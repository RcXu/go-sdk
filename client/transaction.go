package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type BeginTransactionRequest struct {
	StoreName           string
	BunchTransactionNum int32
}

type BeginResponse struct {
	TransactionID       int
	BunchTransactionIDs []string
}

func (c *GRPCClient) DistributeTransactionBegin(ctx context.Context, in *BeginTransactionRequest) (*BeginResponse, error) {
	fmt.Print(in)
	res, err := c.protoClient.DistributeTransactionBegin(ctx, in)
	fmt.Print(res)
	fmt.Print(err)
	if err != nil {
		return &BeginResponse{}, errors.New("transaction begin failed")
	}
	return &BeginResponse{TransactionID: res.TransactionID, BunchTransactionIDs: res.BunchTransactionIDs}, nil
}
