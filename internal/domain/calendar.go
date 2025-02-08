package domain

import (
	"time"

	"gorm.io/gorm"
)

type EventType string

const (
	Event        EventType = "EVENT"
	Meet         EventType = "MEET"
	WorkingHours EventType = "WORKING_HOURS"
)

type CalendarEvent struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description *string   `gorm:"type:text"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Type        EventType `gorm:"type:varchar(20);default:EVENT"`
	CreatorID   string    `gorm:"type:char(24);not null"`
	WorkspaceID string    `gorm:"type:char(24);not null"`
	AssigneeIDs []string  `gorm:"type:char(24)[]"`
	MeetLink    string
}

func (CalendarEvent) TableName() string {
	return "events"
}
