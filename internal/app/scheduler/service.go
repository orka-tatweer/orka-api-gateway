package scheduler

import (
	"github.com/lai0xn/orka/internal/domain"
	"github.com/lai0xn/orka/pkg/client"
)

type TaskService struct {
	repo Repo
}

func NewTaskService(repo Repo) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task TaskDTO) error {
	return s.repo.SaveTask(&domain.Task{
		Name:              task.Name,
		Duration:          task.Duration,
		Dependencies:      task.Dependencies,
		Location:          task.Location,
		Priority:          task.Priority,
		EarliestStart:     task.EarliestStart,
		LatestEnd:         task.LatestEnd,
		CostPerHour:       task.CostPerHour,
		ResourcesRequired: mapResources(task.ResourcesRequired),
	})
}

func (s *TaskService) ScheduleTask(task client.LogisticsRequestDTO) (*client.OptimizeScheduleResponse, error) {
	apiClient := client.NewApiClient()
	response, err := apiClient.ScheduleTask(task)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func mapResources(resources map[string]int) []domain.ResourceRequirement {
	var result []domain.ResourceRequirement
	for resourse, amount := range resources {
		result = append(result, domain.ResourceRequirement{
			Resource: resourse,
			Quantity: amount,
		})
	}
	return result
}

func (s *TaskService) RemoveTask(taskID uint) error {
	return s.repo.DeleteTask(taskID)
}
