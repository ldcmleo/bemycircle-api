package models

import (
	"time"

	"github.com/google/uuid"
)

type Connection struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	ConnectedUserID uuid.UUID `gorm:"type:uuid;not null"`
	Relationship    string    `gorm:"type:varchar(20);not null;check:relationship IN ('friend', 'colleague', 'family', 'partner')"`
	Status          string    `gorm:"type:varchar(20);default:'active';check:status IN ('active', 'former')"`
	IsPrivate       bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	User            User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ConnectedUser   User      `gorm:"foreignKey:ConnectedUserID;constraint:OnDelete:CASCADE"`
}
