package service

import (
	"context"
	"fmt"
	pb "github.com/kvant-node/api/v2/api_pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Service) Status(context.Context, *empty.Empty) (*pb.StatusResponse, error) {
	result, err := s.client.Status()
	if err != nil {
		return new(pb.StatusResponse), status.Error(codes.Internal, err.Error())
	}

	return &pb.StatusResponse{
		Version:           s.version,
		LatestBlockHash:   fmt.Sprintf("%X", result.SyncInfo.LatestBlockHash),
		LatestAppHash:     fmt.Sprintf("%X", result.SyncInfo.LatestAppHash),
		LatestBlockHeight: fmt.Sprintf("%d", result.SyncInfo.LatestBlockHeight),
		LatestBlockTime:   result.SyncInfo.LatestBlockTime.Format(time.RFC3339Nano),
		KeepLastStates:    fmt.Sprintf("%d", s.minterCfg.BaseConfig.KeepLastStates),
		CatchingUp:        result.SyncInfo.CatchingUp,
		PublicKey: &pb.StatusResponse_PubKey{
			//Type:  "todo",
			Value: fmt.Sprintf("Kp%x", result.ValidatorInfo.PubKey.Bytes()[5:]),
		},
	}, nil
}
