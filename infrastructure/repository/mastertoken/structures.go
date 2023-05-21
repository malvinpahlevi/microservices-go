package mastertoken

import (
	"time"
)

type MasterToken struct {
	ID        string    `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Token     string    `json:"token" gorm:"type:text"`
	IssuedAt  time.Time `json:"issued_at, omitempty"`
	ExpiredAt time.Time `json:"expired_at, omitempty"`
}

func (*MasterToken) TableName() string {
	return "master_tokens"
}
