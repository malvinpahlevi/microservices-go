// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"gorm.io/gorm"
	userService "microservices-go/application/usecases/user"
	userRepository "microservices-go/infrastructure/repository/user"
	userController "microservices-go/infrastructure/rest/controllers/user"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := userService.Service{UserRepository: uRepository}
	return &userController.Controller{UserService: service}
}
