// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	medicineController "microservices-go/infrastructure/rest/controllers/medicine"
)

// MedicineRoutes is a function that contains all medicine routes
func MedicineRoutes(router *gin.RouterGroup, controller *medicineController.Controller) {

	routerMedicine := router.Group("/medicine")
	{
		routerMedicine.POST("/", controller.NewMedicine)
		routerMedicine.GET("/:id", controller.GetMedicinesByID)
		routerMedicine.GET("/", controller.GetAllMedicines)
		routerMedicine.PUT("/:id", controller.UpdateMedicine)
		routerMedicine.DELETE("/:id", controller.DeleteMedicine)
	}

}
