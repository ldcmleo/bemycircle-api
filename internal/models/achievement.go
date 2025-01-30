package models

import (
	"time"

	"github.com/google/uuid"
)

type Achievement struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Title       string    `gorm:"not null"`
	Description string
	Category    string    `gorm:"type:varchar(20);not null;check:category IN ('education', 'work', 'personal')"`
	Date        time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
