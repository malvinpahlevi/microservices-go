// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"gorm.io/gorm"
	authService "microservices-go/application/usecases/auth"
	userRepository "microservices-go/infrastructure/repository/user"
	authController "microservices-go/infrastructure/rest/controllers/auth"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *gorm.DB) *authController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := authService.Service{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
