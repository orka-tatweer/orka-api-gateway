package domain

import "gorm.io/gorm"

type ResourceRequirement struct {
	gorm.Model
	TaskID   uint   `json:"task_id"`
	Resource string `json:"resource"`
	Quantity int    `json:"quantity"`
}
