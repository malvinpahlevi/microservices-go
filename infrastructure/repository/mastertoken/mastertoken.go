package mastertoken

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	domainErrors "microservices-go/domain/errors"
	domainMasterToken "microservices-go/domain/mastertoken"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(newMasterToken *domainMasterToken.MasterToken) (createdMasterToken *domainMasterToken.MasterToken, err error) {
	masterToken := fromDomainMapper(newMasterToken)

	tx := r.DB.Create(masterToken)

	if tx.Error != nil {
		fmt.Printf("error %v", tx.Error)
		byteErr, _ := json.Marshal(tx.Error)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return
	}

	createdMasterToken = masterToken.toDomainMapper()
	return
}

func (r *Repository) GetByName(name string) (*domainMasterToken.MasterToken, error) {
	var masterToken MasterToken
	err := r.DB.Where("name = ?", name).First(&masterToken).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return &domainMasterToken.MasterToken{}, err
	}

	return masterToken.toDomainMapper(), nil
}

func (r *Repository) Update(name string, masterTokenMap map[string]interface{}) (*domainMasterToken.MasterToken, error) {
	var masterToken MasterToken

	masterToken.Name = name
	err := r.DB.Model(&masterToken).
		Select("token", "issued_at", "expired_at").
		Updates(masterTokenMap).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainMasterToken.MasterToken{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
			return &domainMasterToken.MasterToken{}, err

		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
			return &domainMasterToken.MasterToken{}, err
		}
	}

	err = r.DB.Where("name = ?", name).First(&masterToken).Error

	return masterToken.toDomainMapper(), err
}
