package client

import (
	"context"
	"fmt"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/pkg/errors"
)

type LogstorageResponse struct {
	Success bool
}

func (c *GRPCClient) OnLogMessage(ctx context.Context, in *pb.LogstorageMessageRequest) (*LogstorageResponse, error) {
	fmt.Print(in)
	res, err := c.protoClient.OnLogMessage(ctx, in)
	fmt.Print(res)
	fmt.Print(err)
	if err != nil {
		return &LogstorageResponse{}, errors.New("request is failed")
	}
	return &LogstorageResponse{Success: res.Success}, nil
}
