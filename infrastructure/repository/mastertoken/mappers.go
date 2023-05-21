package mastertoken

import (
	domainMasterToken "microservices-go/domain/mastertoken"
)

func (masterToken *MasterToken) toDomainMapper() *domainMasterToken.MasterToken {
	return &domainMasterToken.MasterToken{
		ID:        masterToken.ID,
		Name:      masterToken.Name,
		Token:     masterToken.Token,
		IssuedAt:  masterToken.IssuedAt,
		ExpiredAt: masterToken.ExpiredAt,
	}
}

func fromDomainMapper(masterToken *domainMasterToken.MasterToken) *MasterToken {
	return &MasterToken{
		ID:        masterToken.ID,
		Name:      masterToken.Name,
		Token:     masterToken.Token,
		IssuedAt:  masterToken.IssuedAt,
		ExpiredAt: masterToken.ExpiredAt,
	}
}
