package rpc

import (
	"context"

	"github.com/lai0xn/orka/pkg/pb"
)

type Scheduler struct{}

func (s *Scheduler) ScheduleTasks(ctx context.Context, req *pb.LogisticsRequest) (*pb.OptimizeScheduleResponse, error) {
	return nil, nil
}
