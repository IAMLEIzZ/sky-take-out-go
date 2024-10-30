package categorydao

import (
	"sky-take-out-go/db"
	"sky-take-out-go/model/entity"
)

func Save(category *entity.Category) error {
	
	err := db.DB.Debug().Create(category)

	return err.Error 
}