package client

import (
	"context"
	"fmt"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/pkg/errors"
)

type BeginResponse struct {
	TransactionID       int
	BunchTransactionIDs []string
}

func (c *GRPCClient) DistributeTransactionBegin(ctx context.Context, in *pb.BeginTransactionRequest) (*BeginResponse, error) {
	fmt.Print(in)
	res, err := c.protoClient.DistributeTransactionBegin(ctx, in)
	fmt.Print(res)
	fmt.Print(err)
	if err != nil {
		return &BeginResponse{}, errors.New("transaction begin failed")
	}
	return &BeginResponse{TransactionID: res.TransactionID, BunchTransactionIDs: res.BunchTransactionIDs}, nil
}
