package mastertoken

import "time"

type NewMasterToken struct {
	ID        string
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
