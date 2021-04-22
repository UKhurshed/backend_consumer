package service

import (
	"backend_consumer/pkg/domain"
	"backend_consumer/pkg/repository"
)

type BuildingList interface {

}

type BuildingItem interface {
	CreateBuildingItem(building domain.Building) (int, error)
	Delete(buildingId int) error
	GetAll() ([]domain.Building, error)
	Update(buildingId int, building domain.BuildingUpdateInput) error
}

type Service struct {
	BuildingList
	BuildingItem
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		BuildingItem: NewBuildingItemService(repos.BuildingItem),
	}
}