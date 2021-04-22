package domain

import "errors"

type Building struct {
	Id                 int    `json:"-" db:"id"`
	Name               string `json:"name" db:"name" binding:"required"`
	Address            string `json:"address" db:"address" binding:"required"`
	Phone              string `json:"phone" db:"phone" binding:"required"`
	NameBusinessEntity string `json:"name_business_entity" db:"name_business_entity" binding:"required"`
}

type BuildingUpdateInput struct {
	Name               *string `json:"name"`
	Address            *string `json:"address"`
	Phone              *string `json:"phone"`
	NameBusinessEntity *string `json:"name_business_entity"`
}

func (i BuildingUpdateInput) Validate() error {
	if i.Name == nil && i.Phone == nil && i.Address == nil && i.NameBusinessEntity == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
