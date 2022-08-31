package client

import (
	"context"
	"fmt"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"github.com/pkg/errors"
)

type BeginResponse struct {
	TransactionID       string
	BunchTransactionIDs []string
}

type GetDistributeTransactionStateResponse struct {
	TransactionID          string
	BunchTransactionStates map[string]string
}

type DistributeTransactionScheduleResponse struct {
	Message string
}

func (c *GRPCClient) DistributeTransactionBegin(ctx context.Context, in *pb.BeginTransactionRequest) (*BeginResponse, error) {
	fmt.Printf("%v \n", in)
	res, err := c.protoClient.DistributeTransactionBegin(ctx, in)
	fmt.Printf("%v \n", res)
	fmt.Printf("%v \n", err)
	if err != nil {
		return &BeginResponse{}, errors.New("transaction begin failed")
	}
	return &BeginResponse{TransactionID: res.TransactionID, BunchTransactionIDs: res.BunchTransactionIDs}, nil
}

func (c *GRPCClient) GetDistributeTransactionState(ctx context.Context, in *pb.GetDistributeTransactionStateRequest) (*GetDistributeTransactionStateResponse, error) {
	fmt.Printf("%v \n", in)
	res, err := c.protoClient.GetDistributeTransactionState(ctx, in)
	fmt.Printf("%v \n", res)
	fmt.Printf("%v \n", err)
	if err != nil {
		return &GetDistributeTransactionStateResponse{}, errors.New("transaction get state failed")
	}
	return &GetDistributeTransactionStateResponse{TransactionID: res.TransactionID, BunchTransactionStates: res.BunchTransactionStates}, nil
}

func (c *GRPCClient) DistributeTransactionCommit(ctx context.Context, in *pb.DistributeTransactionScheduleRequest) (*DistributeTransactionScheduleResponse, error) {
	fmt.Printf("%v \n", in)
	res, err := c.protoClient.DistributeTransactionCommit(ctx, in)
	fmt.Printf("%v \n", res)
	fmt.Printf("%v \n", err)
	if err != nil {
		return &DistributeTransactionScheduleResponse{}, errors.New("transaction commit failed")
	}
	return &DistributeTransactionScheduleResponse{Message: "transaction commit success"}, nil
}
func (c *GRPCClient) DistributeTransactionRollback(ctx context.Context, in *pb.DistributeTransactionScheduleRequest) (*DistributeTransactionScheduleResponse, error) {
	fmt.Printf("%v \n", in)
	res, err := c.protoClient.DistributeTransactionCommit(ctx, in)
	fmt.Printf("%v \n", res)
	fmt.Printf("%v \n", err)
	if err != nil {
		return &DistributeTransactionScheduleResponse{}, errors.New("transaction rollback failed")
	}
	return &DistributeTransactionScheduleResponse{Message: "transaction rollback success"}, nil
}
