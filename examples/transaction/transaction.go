package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	_ "strconv"

	"gopkg.in/alecthomas/kingpin.v2"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/examples/actor/api"
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
	kingpin.Command("request", "request a distribute transaction")
	kingpin.Command("actor", "request a distribute transaction")

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
	case "request":
		requestBunchTranaction(client)
	case "actor":
		requestActortBunchTranaction(client)
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

func requestBunchTranaction(client dapr.Client) {
	ctx := context.Background()
	data := make(map[string]string)
	data["distribute-transaction-id"] = "transaction-430cb789-066c-4cfe-a98d-5987b0a575b6-6167"
	data["distribute-bunch-transaction-id"] = "bunch-1"
	data["distribute-transaction-store"] = "redis"
	data["say"] = "hello"
	d, _ := json.Marshal(data)
	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        []byte(d),
	}
	resp, err := client.InvokeMethodWithContent(ctx, "serving", "echo", "post", content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("service method invoked, response: %s\n", string(resp))
}

func requestActortBunchTranaction(client dapr.Client) {
	ctx := context.Background()
	myActor := new(api.ClientStub)
	client.ImplActorClientStub(myActor)
	data := make(map[string]string)
	data["distribute-transaction-id"] = "transaction-430cb789-066c-4cfe-a98d-5987b0a575b6-6167"
	data["distribute-bunch-transaction-id"] = "bunch-1"
	data["distribute-transaction-store"] = "redis"
	data["say"] = "hello"
	d, _ := json.Marshal(data)
	rsp, err := myActor.Invoke(ctx, string(d))
	if err != nil {
		panic(err)
	}
	fmt.Println("get invoke result = ", rsp)
}
