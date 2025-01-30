package models

import (
	"time"

	"github.com/google/uuid"
)

type ConnectionRequest struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SenderID     uuid.UUID `gorm:"type:uuid;not null"`
	ReceiverID   uuid.UUID `gorm:"type:uuid;not null"`
	Relationship string    `gorm:"type:varchar(20);not null;check:relationship IN ('friend', 'colleague', 'family', 'partner')"`
	Status       string    `gorm:"type:varchar(20);default:'pending';check:status IN ('pending', 'accepted', 'rejected')"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Sender       User      `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	Receiver     User      `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE"`
}
