package service

import (
	"backend_consumer/pkg/domain"
	"backend_consumer/pkg/repository"
)

type BuildingItemService struct {
	repo repository.BuildingItem
}

func NewBuildingItemService(repo repository.BuildingItem) *BuildingItemService {
	return &BuildingItemService{repo: repo}
}

func (s *BuildingItemService) CreateBuildingItem(building domain.Building) (int, error) {
	return s.repo.CreateBuildingItem(building)
}

func (s *BuildingItemService) Delete(buildingId int) error {
	return s.repo.Delete(buildingId)
}

func (s *BuildingItemService) GetAll(nameBuilding, typeOfObject, networkTrading, region string) ([]domain.BuildingSelect, error) {
	return s.repo.GetAll(nameBuilding, typeOfObject, networkTrading, region)
}

func (s *BuildingItemService) Update(buildingId int, building domain.BuildingUpdateInput) error {
	if err := building.Validate(); err != nil {
		return err
	}
	return s.repo.Update(buildingId, building)
}
