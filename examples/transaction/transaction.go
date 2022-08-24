package main

import (
	"context"
	"fmt"
	"os"
	_ "strconv"

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

	// create the client
	client, err := dapr.NewClientWithPort(port)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	ctx := context.Background()

	logRequest := pb.BeginTransactionRequest{
		StoreName:           "redis",
		BunchTransactionNum: 2,
	}
	reqs, err := client.DistributeTransactionBegin(ctx, &logRequest)
	if err != nil {
		fmt.Printf("Failed to begin a transaction: %v\n", err)
	}
	fmt.Print(reqs)
}
