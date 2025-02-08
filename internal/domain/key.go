package domain

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	Key        string
	WebHookUrl string
	User       User
	UserID     uint
}
