package rpc

import (
	"context"

	"github.com/lai0xn/orka/pkg/client"
	"github.com/lai0xn/orka/pkg/pb"
	"github.com/twitchtv/twirp"
)

type Scheduler struct {
	apiClient client.APIClient
}

func (s *Scheduler) ScheduleTasks(ctx context.Context, req *pb.LogisticsRequest) (*pb.OptimizeScheduleResponse, error) {
	var tasks []client.TaskDTO
	for _, t := range req.Tasks {
		tasks = append(tasks, client.TaskDTO{
			ID:                int(t.Id),
			Name:              t.Name,
			Duration:          int(t.Duration),
			Dependencies:      convertUint32Slice(t.Dependencies),
			ResourcesRequired: convertResourceMap(t.ResourcesRequired),
			Location:          t.Location,
			Priority:          int(t.Priority),
			EarliestStart:     int(t.EarliestStart),
			LatestEnd:         int(t.LatestEnd),
			CostPerHour:       float64(t.CostPerHour),
		})
	}

	// Call API
	resp, err := s.apiClient.ScheduleTask(client.LogisticsRequestDTO{
		Tasks:         tasks,
		Vehicles:      req.Vehicles,
		Objective:     req.Objective,
		ResourcePool:  req.ResourcePool,
		TransitMatrix: convertTransitMap(req.TransitMatrix),
	})
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}

	// Convert response back to pb.OptimizeScheduleResponse
	return &pb.OptimizeScheduleResponse{
		Result: &pb.ScheduleResponse{
			Schedule:  convertScheduleMap(resp.Result.Schedule),
			Makespan:  convertOptionalInt(resp.Result.Makespan),
			TotalCost: convertOptionalFloat(resp.Result.TotalCost),
		},
	}, nil
}

func convertUint32Slice(dependencies []uint32) []uint {
	var result []uint
	for _, d := range dependencies {
		result = append(result, uint(d))
	}
	return result
}

func convertResourceMap(resources map[string]int32) map[string]int {
	result := make(map[string]int)
	for k, v := range resources {
		result[k] = int(v)
	}
	return result
}

func convertOptionalInt32(val *int32) *int {
	if val == nil {
		return nil
	}
	v := int(*val)
	return &v
}

func convertOptionalFloat32(val *float32) *float64 {
	if val == nil {
		return nil
	}
	v := float64(*val)
	return &v
}

func convertTransitMap(pbTransit map[string]*pb.TransitMap) map[string]map[string]int {
	result := make(map[string]map[string]int)
	for key, tm := range pbTransit {
		result[key] = make(map[string]int)
		for k, v := range tm.TransitTimes {
			result[key][k] = int(v)
		}
	}
	return result
}

func convertScheduleMap(schedule map[string]client.ScheduleTask) map[string]*pb.ScheduleTask {
	result := make(map[string]*pb.ScheduleTask)
	for key, task := range schedule {
		result[key] = &pb.ScheduleTask{
			Name:      task.Name,
			Start:     int32(task.Start),
			End:       int32(task.End),
			Resources: convertResourceMapReverse(task.Resources),
			Location:  task.Location,
			Vehicle:   task.Vehicle,
		}
	}
	return result
}

func convertResourceMapReverse(resources map[string]int) map[string]int32 {
	result := make(map[string]int32)
	for k, v := range resources {
		result[k] = int32(v)
	}
	return result
}

func convertOptionalInt(val *int) *int32 {
	if val == nil {
		return nil
	}
	v := int32(*val)
	return &v
}

func convertOptionalFloat(val *float64) *float32 {
	if val == nil {
		return nil
	}
	v := float32(*val)
	return &v
}

func (s *Scheduler) GenerateProductionPlan(context.Context, *pb.ProductionPlanRequest) (*pb.ProductionResponse, error) {
	return nil, nil
}
