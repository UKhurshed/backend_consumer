package repository

import (
	"backend_consumer/pkg/domain"
	"github.com/jmoiron/sqlx"
)

type BuildingList interface {

}

type BuildingItem interface {
	CreateBuildingItem(building domain.Building) (int, error)
	Delete(buildingId int) error
	GetAll() ([]domain.Building, error)
	Update(buildingId int, building domain.BuildingUpdateInput) error
}

type Repository struct {
	BuildingList
	BuildingItem
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		BuildingItem: NewBuildingPostgres(db),
	}
}
