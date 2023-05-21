package mastertoken

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	useCaseMasterToken "microservices-go/application/usecases/mastertoken"
	domainError "microservices-go/domain/errors"
	domainMasterToken "microservices-go/domain/mastertoken"
	"microservices-go/infrastructure/rest/controllers"
	"net/http"
	"time"
)

type Controller struct {
	MasterTokenService useCaseMasterToken.Service
}

func (c *Controller) NewMasterToken(ctx *gin.Context) {
	var request NewMasterTokenRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		fmt.Printf("error %v", err)
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newMedicine := useCaseMasterToken.NewMasterToken{
		ID:        uuid.New().String(),
		Name:      request.Name,
		Token:     request.Token,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(50 * time.Minute),
	}

	domainMasterToken, err := c.MasterTokenService.Create(&newMedicine)
	if err != nil {
		fmt.Printf("error %v", err)
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainMasterToken)

}

func (c *Controller) GetMasterTokenByName(ctx *gin.Context) {
	name := ctx.Query("name")

	domainMasterToken, err := c.MasterTokenService.GetByName(name)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainMasterToken)
}

func (c *Controller) UpdateMasterToken(ctx *gin.Context) {
	var err error

	name, ok := ctx.GetQuery("name")
	if !ok {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	var requestMap map[string]interface{}

	switch name {
	case "gdrive":
		requestMap["issued_at"] = time.Now()
		requestMap["issued_at"] = time.Now().Add(50 * time.Minute)
	}

	err = controllers.BindJSONMap(ctx, &requestMap)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = updateValidation(requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var masterToken *domainMasterToken.MasterToken
	masterToken, err = c.MasterTokenService.Update(name, requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, masterToken)

}
