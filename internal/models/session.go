package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID	uuid.UUID	`gorm:"type:uuid;primaryKey"`
	CampaignID	uuid.UUID	`gorm:"type:uuid;index;not null"`
	CreatedAt	time.Time	`gorm:"not null"`
}