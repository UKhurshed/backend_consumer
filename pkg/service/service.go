package service

import (
	"backend_consumer/pkg/domain"
	"backend_consumer/pkg/repository"
)

type BuildingItem interface {
	CreateBuildingItem(building domain.Building) (int, error)
	Delete(buildingId int) error
	GetAll() ([]domain.Building, error)
	Update(buildingId int, building domain.BuildingUpdateInput) error
}

type SubjectItem interface {
	GetAllSubjects() ([]domain.Subject, error)
	CreateSubject(subject domain.Subject) (int, error)
	UpdateSubject(subjectId int, subject domain.SubjectInput) error
	DeleteSubject(subjectId int) error
}

type Service struct {
	BuildingItem
	SubjectItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		BuildingItem: NewBuildingItemService(repos.BuildingItem),
		SubjectItem:  NewSubjectService(repos.Subject),
	}
}
