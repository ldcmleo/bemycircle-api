package models

import (
	"time"

	"github.com/google/uuid"
)

type PrivacySetting struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	ShowFriends     bool      `gorm:"default:true"`
	ShowColleagues  bool      `gorm:"default:true"`
	ShowFamily      bool      `gorm:"default:true"`
	ShowPartners    bool      `gorm:"default:true"`
	ShowAchivements bool      `gorm:"default:true"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	User            User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
