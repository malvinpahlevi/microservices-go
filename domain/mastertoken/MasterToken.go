package mastertoken

import (
	"time"
)

type MasterToken struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	IssuedAt  time.Time `json:"issued_at, omitempty"`
	ExpiredAt time.Time `json:"expired_at, omitempty"`
}

type Service interface {
	Get(string) (*MasterToken, error)
	Create(*MasterToken) error
	GetByName(string) (*MasterToken, error)
	Update(string, map[string]interface{}) (*MasterToken, error)
}
