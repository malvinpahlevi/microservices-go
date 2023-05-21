package mastertoken

import (
	domainMasterToken "microservices-go/domain/mastertoken"
)

func (n *NewMasterToken) toDomainMapper() *domainMasterToken.MasterToken {
	return &domainMasterToken.MasterToken{
		ID:        n.ID,
		Name:      n.Name,
		Token:     n.Token,
		IssuedAt:  n.IssuedAt,
		ExpiredAt: n.ExpiredAt,
	}
}
