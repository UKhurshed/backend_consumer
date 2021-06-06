package repository

import (
	"backend_consumer/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type BuildingItem interface {
	CreateBuildingItem(building domain.Building) (int, error)
	Delete(buildingId int) error
	GetAll(nameBuilding, typeOfObject, networkTrading, region string) ([]domain.BuildingSelect, error)
	Update(buildingId int, building domain.BuildingUpdateInput) error
}

type Subject interface {
	GetAllSubjects() ([]domain.Subject, error)
	CreateSubject(subject domain.Subject) (int, error)
	UpdateSubject(subjectId int, subject domain.SubjectInput) error
	DeleteSubject(subjectId int) error
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(email, password string) (domain.User, error)
}

type Repository struct {
	BuildingItem
	Subject
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		BuildingItem: NewBuildingPostgres(db),
		Subject:      NewSubjectPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
