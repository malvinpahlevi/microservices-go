// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"gorm.io/gorm"
	medicineService "microservices-go/application/usecases/medicine"
	medicineRepository "microservices-go/infrastructure/repository/medicine"
	medicineController "microservices-go/infrastructure/rest/controllers/medicine"
)

// MedicineAdapter is a function that returns a medicine controller
func MedicineAdapter(db *gorm.DB) *medicineController.Controller {
	mRepository := medicineRepository.Repository{DB: db}
	service := medicineService.Service{MedicineRepository: mRepository}
	return &medicineController.Controller{MedicineService: service}
}
