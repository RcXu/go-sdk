package main

import (
	"context"
	"fmt"
	"os"
	_ "strconv"

	"gopkg.in/alecthomas/kingpin.v2"
	_ "gopkg.in/alecthomas/kingpin.v2"

	dapr "github.com/dapr/go-sdk/client"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
)

const (
	daprPort = "3500"
)

var port string

func init() {
	if port = os.Getenv("DAPR_GRPC_PORT"); len(port) == 0 {
		port = daprPort
	}
}

func main() {
	kingpin.Command("begin", "Begin a distribute transaction")
	Get := kingpin.Command("get", "Get distribute transaction state.")
	var transactionID string
	Get.Flag("id", "transaction ID.").Default("1").StringVar(&transactionID)

	// create the client
	client, err := dapr.NewClientWithPort(port)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	switch kingpin.Parse() {
	case "begin":
		beginAction(client)
	case "get":
		getStateAction(client, transactionID)
	}

}

func beginAction(client dapr.Client) {
	ctx := context.Background()
	request := pb.BeginTransactionRequest{
		StoreName:           "redis",
		BunchTransactionNum: 2,
	}
	reqs, err := client.DistributeTransactionBegin(ctx, &request)
	if err != nil {
		fmt.Printf("Failed to begin a transaction: %v\n", err)
	}
	fmt.Print(reqs)
}

func getStateAction(client dapr.Client, transactionID string) {
	ctx := context.Background()
	request := pb.GetDistributeTransactionStateRequest{
		StoreName:     "redis",
		TransactionID: transactionID,
	}
	fmt.Printf("request is : %v\n", request)
	reqs, err := client.GetDistributeTransactionState(ctx, &request)

	if err != nil {
		fmt.Printf("Failed to get transaction state : %v\n", err)
	}
	fmt.Printf("state is : %v\n", reqs)
}
