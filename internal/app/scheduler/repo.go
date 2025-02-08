package scheduler

import "github.com/lai0xn/orka/internal/domain"

type Repo interface {
	SaveTask(task *domain.Task) error
	DeleteTask(taskID uint) error
}
