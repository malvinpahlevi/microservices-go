package adapter

import (
	"gorm.io/gorm"
	masterTokenService "microservices-go/application/usecases/mastertoken"
	masterTokenRepository "microservices-go/infrastructure/repository/mastertoken"
	masterTokenController "microservices-go/infrastructure/rest/controllers/mastertoken"
)

func MasterTokenAdapter(db *gorm.DB) *masterTokenController.Controller {
	mRepository := masterTokenRepository.Repository{DB: db}
	service := masterTokenService.Service{MasterTokenRepository: mRepository}
	return &masterTokenController.Controller{MasterTokenService: service}
}
