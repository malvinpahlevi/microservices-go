package mastertoken

import (
	masterTokenDomain "microservices-go/domain/mastertoken"
)
import masterTokenRepository "microservices-go/infrastructure/repository/mastertoken"

type Service struct {
	MasterTokenRepository masterTokenRepository.Repository
}

func (s *Service) Create(masterToken *NewMasterToken) (*masterTokenDomain.MasterToken, error) {
	masterTokenModel := masterToken.toDomainMapper()
	return s.MasterTokenRepository.Create(masterTokenModel)
}

func (s *Service) GetByName(name string) (*masterTokenDomain.MasterToken, error) {
	return s.MasterTokenRepository.GetByName(name)
}

// Update is a function that updates a medicine by id
func (s *Service) Update(name string, medicineMap map[string]interface{}) (*masterTokenDomain.MasterToken, error) {
	return s.MasterTokenRepository.Update(name, medicineMap)
}
