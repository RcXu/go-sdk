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
	var getTransactionID string
	Get.Flag("id", "transaction ID.").Default("1").StringVar(&getTransactionID)

	var reqTransactionID string
	Request := kingpin.Command("request", "request a distribute transaction")
	Request.Flag("id", "transaction ID.").Default("1").StringVar(&reqTransactionID)

	kingpin.Command("actor", "request a distribute transaction")

	var comTransactionID string
	Commit := kingpin.Command("commit", "request a distribute transaction")
	Commit.Flag("id", "transaction ID.").Default("1").StringVar(&comTransactionID)

	kingpin.Command("rollback", "request a distribute transaction")

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
		getStateAction(client, getTransactionID)
	case "request":

		requestBunchTranaction(client, reqTransactionID)
	case "actor":
		requestActortBunchTranaction(client)
	case "commit":
		requestCommitTransaction(client, comTransactionID)
	case "rollback":
		requestRollbackTransaction(client)
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

func requestBunchTranaction(client dapr.Client, transactionID string) {
	ctx := context.Background()
	data := make(map[string]string)
	data["distribute-transaction-id"] = transactionID
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
	data["distribute-transaction-id"] = "transaction-efda9c9f-429e-4369-981d-db3de20c8b58-9108"
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

func requestCommitTransaction(client dapr.Client, transactionID string) {
	ctx := context.Background()
	request := pb.DistributeTransactionScheduleRequest{
		StoreName:     "redis",
		TransactionID: transactionID,
	}
	fmt.Printf("request is : %v\n", request)
	reqs, err := client.DistributeTransactionCommit(ctx, &request)

	if err != nil {
		fmt.Printf("Failed to commit transaction  : %v\n", err)
	}
	fmt.Print(reqs)
}

func requestRollbackTransaction(client dapr.Client) {
	ctx := context.Background()
	request := pb.DistributeTransactionScheduleRequest{
		StoreName:     "redis",
		TransactionID: "transaction-430cb789-066c-4cfe-a98d-5987b0a575b6-6167",
	}
	fmt.Printf("request is : %v\n", request)
	reqs, err := client.DistributeTransactionRollback(ctx, &request)

	if err != nil {
		fmt.Printf("Failed to commit transaction  : %v\n", err)
	}
	fmt.Print(reqs)
}
