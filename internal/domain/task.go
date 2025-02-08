package domain

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name              string                `json:"name"`
	Duration          int                   `json:"duration"`
	Location          string                `json:"location"`
	Priority          int                   `json:"priority"`
	EarliestStart     *int                  `json:"earliest_start"`
	LatestEnd         *int                  `json:"latest_end"`
	CostPerHour       *float64              `json:"cost_per_hour"`
	Dependencies      []uint                `json:"dependencies" gorm:"type:json"`
	ResourcesRequired []ResourceRequirement `json:"resources_required" gorm:"foreignKey:TaskID"`
}
