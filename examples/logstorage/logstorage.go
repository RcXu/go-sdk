/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	dapr "github.com/dapr/go-sdk/client"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
)

const (
	logstorageName = `alicloud`
	daprPort       = "3500"
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

	item, err := client.OnLogMessage(ctx, &pb.LogstorageMessageRequest{
		LogstorageName: "alicloud",
		Project:"k8s-log-c134d574cd3e6405682e2dda2095ea35d",
		Logstore:"person-log",
		Topic:"dapr-demo",
		Source:"dapr",
		Log: &pb.LogstorageMessageContent{
			Ip:"127.0.0.1",
			Timestamp:time.Now().Unix(),
			File:"/demo/logstorage/api.go",
			Function:"OnLogMessage",
			Level:"debug",
			Content:"this is dapr grpc log demo",
		},
	})
	if err != nil {
		fmt.Printf("Failed to log message: %v\n", err)
	}
	fmt.Printf("Log Message Success\n")
}
