package client

import (
	"context"
	"fmt"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/pkg/errors"
)

func (c *GRPCClient) OnLogMessage(ctx context.Context, in *pb.LogstorageMessageRequest) error {
	fmt.Print(in)
	res, err := c.protoClient.OnLogMessage(ctx, in)
	fmt.Print(res)
	fmt.Print(err)
	return errors.New("request is failed")

}
