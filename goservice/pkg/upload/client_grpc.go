package upload

import (
	"context"

	pb "github.com/RicardoSandoval11/MediaPipeline/goservice/pkg/pb/file_counter_service"
	"github.com/go-kit/kit/endpoint"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type UpsertFileCounterRequest struct {
	UserId uuid.UUID
}

type UpsertFileCounterResponse struct {
	Success      bool
	LimitReached bool
	Error        string
}

func NewFileCounterGRPCClient(conn *grpc.ClientConn) endpoint.Endpoint {
	return gt.NewClient(
		conn,
		"file_counter.FileCounterService", // package + service implemented by .NET gRPC server
		"UpsertFileCounterAsync",       // RPC method
		encodeUpsertRequest,            // go struct to proto struct
		decodeUpsertResponse,           // proto struct to go struct
		pb.UpsertFileCounterResponse{}, // expected response struct by protobuf
	).Endpoint()
}

func encodeUpsertRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(UpsertFileCounterRequest)
	return &pb.UpsertFileCounterRequest{
		UserId: req.UserId.String(),
	}, nil
}

func decodeUpsertResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	reply := grpcResp.(*pb.UpsertFileCounterResponse)
	return UpsertFileCounterResponse{
		Success:      reply.GetSuccess(),
		LimitReached: reply.GetLimitReached(),
		Error:        reply.GetError(),
	}, nil
}
