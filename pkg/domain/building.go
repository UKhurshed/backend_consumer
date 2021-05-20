package domain

import (
	"errors"
)

type Building struct {
	Id                 int    `json:"-" db:"id"`
	Name               string `json:"name" db:"name" binding:"required"`
	Address            string `json:"address" db:"address" binding:"required"`
	Phone              string `json:"phone" db:"phone" binding:"required"`
	NameBusinessEntity string `json:"name_business_entity" db:"name_business_entity" binding:"required"`
}

type BuildingSelect struct {
	Id                int     `json:"id" db:"id"`
	NameBuilding      string  `json:"name_building" db:"name_building"`
	NameFullBuilding  string  `json:"name_full_building" db:"name_full_building"`
	ObjectType        bool    `json:"object_type" db:"object_type"`
	SelfService       bool    `json:"self_service" db:"self_service"`
	AvailabilityAsu   bool    `json:"availability_asu" db:"availability_asu"`
	TotalArea         int     `json:"total_area" db:"total_area"`
	RetailSpace       int     `json:"retail_space" db:"retail_space"`
	OpeningDate       *string `json:"opening_date" db:"opening_date"`
	ClosingDate       *string `json:"closing_date" db:"closing_date"`
	WorkPlaceCount    int     `json:"work_place_count" db:"workplacecount"`
	EmployeeCount     int     `json:"employee_count" db:"employee_count"`
	Inn               string  `json:"inn" db:"inn"`
	Kpp               string  `json:"kpp" db:"kpp"`
	TypeObject        string  `json:"type_object" db:"type_object"`
	NetworkTrading    string  `json:"network_trading" db:"network_trading"`
	NameRegion        string  `json:"name_region" db:"name_region"`
	StreetName        *string `json:"street_name" db:"street_name"`
	MicroDistrictName string  `json:"micro_district_name" db:"micro_district_name"`
	FormName          string  `json:"form_name" db:"form_name"`
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
