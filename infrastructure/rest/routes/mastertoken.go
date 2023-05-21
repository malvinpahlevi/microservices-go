// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	masterTokenController "microservices-go/infrastructure/rest/controllers/mastertoken"
)

func MasterTokenRoutes(router *gin.RouterGroup, controller *masterTokenController.Controller) {

	routerMedicine := router.Group("/masterToken")
	{
		routerMedicine.POST("/", controller.NewMasterToken)
		routerMedicine.GET("/", controller.GetMasterTokenByName) // ?name=googleDrive
		routerMedicine.PUT("/:name", controller.UpdateMasterToken)
	}

}
