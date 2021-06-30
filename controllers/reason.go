package controllers

import (
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"wallet/models"
)

type ReasonController struct {
	Db *gorm.DB
}


func (reasonController *ReasonController) Create( reasonObj map[string]interface{} ) error{
	var reason models.Reason

	//Decode JavaScript Object
	err := mapstructure.Decode( reasonObj, &reason )
	if err != nil{
		return err
	}

	val := validator.New()
	err = val.Struct(&reason)
	if err != nil{
		return err
	}

	resp := reasonController.Db.Create( &reason )
	if resp.Error != nil{
		return resp.Error
	}
	return nil
}

func (reasonController *ReasonController) List() ( []models.Reason, error ){
	var reasons []models.Reason
	resp := reasonController.Db.Find( reasons )
	if resp.Error != nil{
		return reasons, resp.Error
	}
	return reasons, nil
}

func (reasonController *ReasonController) Delete( id uint ) error {

	//Decode JavaScript Object
	resp := reasonController.Db.Delete(&models.Reason{}, id )
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

