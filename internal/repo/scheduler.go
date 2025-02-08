package repo

import (
	"errors"

	"github.com/lai0xn/orka/internal/domain"
	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}

func (r *authRepo) SaveTask(task *domain.Task) error {
	if task.ID > 0 {
		return r.db.Save(task).Error
	}
	return r.db.Create(task).Error
}

func (r *authRepo) DeleteTask(taskID uint) error {
	var task domain.Task
	if err := r.db.First(&task, taskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}
	return r.db.Delete(&task).Error
}
