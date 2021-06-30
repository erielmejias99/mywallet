package transaction

import (
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type ReasonRepository struct {
	db * gorm.DB
}

var reasonRepository * ReasonRepository = nil
func GetReasonRepository( db * gorm.DB) * ReasonRepository {
	if reasonRepository == nil{
		reasonRepository = &ReasonRepository{ db: db }
	}
	return reasonRepository
}

func (reasonRep * ReasonRepository) Create( reasonObj map[string]interface{} ) error{
	var reason Reason

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

	resp := reasonRep.db.Create( &reason )
	if resp.Error != nil{
		return resp.Error
	}
	return nil
}

func (reasonRep * ReasonRepository ) List() ( []Reason, error ){
	var reasons []Reason
	resp := reasonRepository.db.Find( reasons )
	if resp.Error != nil{
		return reasons, resp.Error
	}
	return reasons, nil
}

func (reasonRep * ReasonRepository) Delete( id uint ) error {

	//Decode JavaScript Object
	resp := reasonRep.db.Delete(&Reason{}, id )
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
