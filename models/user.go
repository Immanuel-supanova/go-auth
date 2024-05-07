package models

import "time"

type User struct {
	ID        uint
	Email     string `gorm:"unique"`
	Password  string
	IsActive  bool `gorm:"default:true"`
	IsAdmin   bool
	IsStaff   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	LogEntry  LogEntry
}
