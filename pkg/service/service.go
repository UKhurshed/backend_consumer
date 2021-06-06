package service

import (
	"backend_consumer/pkg/domain"
	"backend_consumer/pkg/repository"
)

type BuildingItem interface {
	CreateBuildingItem(building domain.Building) (int, error)
	Delete(buildingId int) error
	GetAll(nameBuilding, typeOfObject, networkTrading, region string) ([]domain.BuildingSelect, error)
	Update(buildingId int, building domain.BuildingUpdateInput) error
}

type SubjectItem interface {
	GetAllSubjects() ([]domain.Subject, error)
	CreateSubject(subject domain.Subject) (int, error)
	UpdateSubject(subjectId int, subject domain.SubjectInput) error
	DeleteSubject(subjectId int) error
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(email string, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Service struct {
	BuildingItem
	SubjectItem
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		BuildingItem: NewBuildingItemService(repos.BuildingItem),
		SubjectItem:  NewSubjectService(repos.Subject),
		Authorization : NewAuthService(repos.Authorization),
	}
}
