package domain

import "errors"

type Building struct {
	Id                int    `json:"-" db:"id"`
	NameBuilding      string `json:"name_building" db:"name_building" binding:"required"`
	NameFullBuilding  string `json:"name_full_buildings" db:"name_full_building" binding:"required"`
	ObjectType        bool   `json:"object_type" db:"object_type" binding:"required"`
	SelfService       bool   `json:"self_service" db:"self_service" binding:"required"`
	AvailabilityAsu   bool   `json:"availability_asu" db:"availability_asu" binding:"required"`
	TotalArea         int    `json:"total_area" db:"total_area" binding:"required"`
	RetailSpace       int    `json:"retail_space" db:"retail_space" binding:"required"`
	OpeningDate       string `json:"opening_date" db:"opening_date" binding:"required"`
	WorkPlaceCount    int    `json:"work_place_count" db:"workplacecount" binding:"required"`
	EmployeeCount     int    `json:"employee_count" db:"employee_count" binding:"required"`
	StreetName        string `json:"street_name" db:"street_name" binding:"required"`
	MicroDistrictName string `json:"micro_district_name" db:"micro_district_name" binding:"required"`
	Inn               string `json:"inn" db:"inn" binding:"required"`
	Kpp               string `json:"kpp" db:"kpp" binding:"required"`
	TypeObjectId      int    `json:"type_object_id" db:"typeOfObject_id" binding:"required"`
	RegionId          int    `json:"region_id" db:"region_id" binding:"required"`
	TradingNetworkId  int    `json:"trading_network_id" db:"tradingNetwork_id" binding:"required"`
	FormOwnerId       int    `json:"form_owner_id" db:"form_owner_id" binding:"required"`
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
	StreetName        string  `json:"street_name" db:"street_name"`
	MicroDistrictName string  `json:"micro_district_name" db:"micro_district_name"`
	FormName          string  `json:"form_name" db:"form_name"`
}

type BuildingUpdateInput struct {
	NameBuilding      *string `json:"name_building"`
	NameFullBuilding  *string `json:"name_full_building"`
	ObjectType        *bool   `json:"object_type"`
	SelfService       *bool   `json:"self_service"`
	AvailabilityAsu   *bool   `json:"availability_asu"`
	TotalArea         *int    `json:"total_area"`
	RetailSpace       *int    `json:"retail_space"`
	OpeningDate       *int    `json:"opening_date"`
	ClosingDate       *string `json:"closing_date"`
	WorkPlaceCount    *int    `json:"work_place_count"`
	EmployeeCount     *int    `json:"employee_count"`
	StreetName        *string `json:"street_name"`
	MicroDistrictName *string `json:"micro_district_name"`
	Inn               *string `json:"inn"`
	Kpp               *string `json:"kpp"`
	RegionId          *int    `json:"region_id"`
	TypeObjectId      *int    `json:"type_object_id"`
	TradingNetworkId  *int    `json:"trading_network_id"`
	FormOwnerId       *int    `json:"form_owner_id"`
}

func (i BuildingUpdateInput) Validate() error {
	if i.NameBuilding == nil && i.NameFullBuilding == nil && i.ObjectType == nil && i.SelfService == nil && i.AvailabilityAsu == nil && i.TotalArea == nil && i.RetailSpace == nil && i.OpeningDate == nil && i.WorkPlaceCount == nil && i.EmployeeCount == nil && i.StreetName == nil && i.MicroDistrictName == nil && i.Inn == nil && i.Kpp == nil && i.RegionId == nil && i.TradingNetworkId == nil && i.TypeObjectId == nil && i.FormOwnerId == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
