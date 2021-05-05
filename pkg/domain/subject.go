package domain

type Subject struct {
	Id                int    `json:"id" db:"id"`
	SubjectName       string `json:"subject_name" db:"subject_name" binding:"required"`
	FullNameSubject   string `json:"full_name_subject" db:"full_name_subject" binding:"required"`
	INN               string `json:"inn" db:"inn" binding:"required"`
	KPP               string `json:"kpp" db:"kpp" binding:"required"`
	FormOfOwnerShipId int    `json:"form_of_owner_ship_id" db:"form_of_ownership_id" binding:"required"`
}

type SubjectInput struct {
	SubjectName       *string `json:"subject_name"`
	FullNameSubject   *string `json:"full_name_subject"`
	INN               *string `json:"inn"`
	KPP               *string `json:"kpp"`
	FormOfOwnerShipId *int    `json:"form_of_owner_ship_id"`
}


