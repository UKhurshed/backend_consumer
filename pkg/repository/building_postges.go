package repository

import (
	"backend_consumer/pkg/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type BuildingPostgres struct {
	db *sqlx.DB
}

func NewBuildingPostgres(db *sqlx.DB) *BuildingPostgres{
	return &BuildingPostgres{
		db: db,
	}
}

func (r *BuildingPostgres) CreateBuildingItem(building domain.Building) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, address, phone, name_business_entity) values ($1, $2, $3, $4) RETURNING id", "buildings")
	row := r.db.QueryRow(query, building.Name, building.Address, building.Phone, building.NameBusinessEntity)

	if err := row.Scan(&id); err!= nil{
		return 0, err
	}

	return id, nil
}

func (r *BuildingPostgres) GetAll() ([]domain.Building, error){
	var items []domain.Building

	query := fmt.Sprintf(`SELECT * FROM buildings`)

	if err := r.db.Select(&items, query); err != nil{
		return nil, err
	}
	return items, nil
}

func (r *BuildingPostgres) Delete(buildingId int) error{
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, "buildings")
	_, err := r.db.Exec(query, buildingId)
	return err
}

func (r *BuildingPostgres) Update(buildingId int, building domain.BuildingUpdateInput) error {
	setValues := make([]string, 0)
	args:= make([]interface{}, 0)
	argId := 1

	if building.Name != nil{
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args  = append(args, *building.Name)
		argId++
	}

	if building.Phone != nil{
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args  = append(args, *building.Phone)
		argId++
	}

	if building.Address != nil{
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args  = append(args, *building.Address)
		argId++
	}

	if building.NameBusinessEntity != nil{
		setValues = append(setValues, fmt.Sprintf("nameBusinessEntity=$%d", argId))
		args  = append(args, *building.NameBusinessEntity)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE buildings SET %s WHERE id=$%d`, setQuery, argId)

	args = append(args, buildingId)

	_, err := r.db.Exec(query, args...)

	return err

}


