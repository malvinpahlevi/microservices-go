package mastertoken

type NewMasterTokenRequest struct {
	Name      string `json:"name" binding:"required"`
	Token     string `json:"token" binding:"required"`
	IssuedAt  string `json:"issued_at"`
	ExpiredAt string `json:"expired_at"`
}
