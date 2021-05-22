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

func NewBuildingPostgres(db *sqlx.DB) *BuildingPostgres {
	return &BuildingPostgres{
		db: db,
	}
}

func (r *BuildingPostgres) CreateBuildingItem(building domain.Building) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name_building, name_full_building, object_type, self_service," +
		"availability_asu, total_area, retail_space, opening_date, workPlaceCount, employee_count, street_name," +
		"micro_district_name, inn, kpp, region_id, typeOfObject_id, tradingNetwork_id, form_owner_id) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18 ) RETURNING id", "BuildingEntity")
	row := r.db.QueryRow(query, building.NameBuilding, building.NameFullBuilding, building.ObjectType, building.SelfService,
		building.AvailabilityAsu, building.TotalArea, building.RetailSpace, building.OpeningDate, building.WorkPlaceCount, building.EmployeeCount,
		building.StreetName, building.MicroDistrictName, building.Inn, building.Kpp, building.RegionId, building.TypeObjectId, building.TradingNetworkId, building.FormOwnerId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BuildingPostgres) GetAll(nameBuilding, typeOfObject, networkTrading, region, microDistrict, streetName, openIn string) ([]domain.BuildingSelect, error) {
	var items []domain.BuildingSelect

	//SELECT
	//       bt.id, bt.name_building, bt.object_type, bt.self_service, bt.availability_asu, bt.total_area,
	//       bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count,
	//       tf.object_type, tn.network_trading, rg.name_region, bt.street_name, bt.micro_district_name
	//       FROM BuildingEntity bt, TypeOfObject tf, TradingNetwork tn, Region rg
	//WHERE bt.region_id = rg.id and bt.typeOfObject_id = tf.id and bt.tradingNetwork_id = tn.id;

	if nameBuilding == "" && microDistrict == "" && streetName == "" && openIn == "" {
		query := fmt.Sprintf(`SELECT bt.id, bt.name_building, bt.name_full_building, bt.object_type, bt.self_service, bt.availability_asu, bt.total_area,
        bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count, bt.inn, bt.kpp,
        tf.type_object, tn.network_trading, rg.name_region, bt.street_name, bt.micro_district_name, fs.form_name
		FROM BuildingEntity bt inner join TypeOfObject tf on bt.typeOfObject_id=tf.id inner join
    	TradingNetwork tn on bt.tradingNetwork_id=tn.id inner join Region rg on bt.region_id = rg.id inner join FormOfOwnerShip fs on bt.id=fs.id;`)
		if err := r.db.Select(&items, query); err != nil {
			return nil, err
		}
	} else {
		query := fmt.Sprintf(`SELECT bt.id, bt.name_building, bt.object_type, bt.self_service, bt.availability_asu, bt.total_area,
        bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count,
        tf.type_object, tn.network_trading, rg.name_region, bt.street_name, bt.micro_district_name
		FROM %s bt inner join %s tf on bt.typeOfObject_id=tf.id inner join
    	%s tn on bt.tradingNetwork_id=tn.id inner join %s rg on bt.region_id = rg.id
		WHERE (bt.name_building = $1 OR bt.region_id = $2 or bt.typeOfObject_id = $3 or
		bt.tradingNetwork_id = $4 or bt.street_name = $5 or bt.micro_district_name = $6) and bt.opening_date > $7`, buildingEntityTable, typeOfObjectsTable, tradingNetworkTable, regionTable)

		//or bt.typeOfObject_id=$2 or bt.tradingNetwork_id=$3 or bt.region_id=$4 or bt.street_name=$5 or bt.micro_district_name > $6
		fmt.Printf("Data" + openIn)
		if err := r.db.Select(&items, query, nameBuilding, region, typeOfObject, networkTrading, streetName, microDistrict, "2000-12-31"); err != nil {
			return nil, err
		}
	}

	return items, nil

	//, typeOfObject, networkTrading, region, streetName, microDistrict
}

func (r *BuildingPostgres) Delete(buildingId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, "buildings")
	_, err := r.db.Exec(query, buildingId)
	return err
}

func (r *BuildingPostgres) Update(buildingId int, building domain.BuildingUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if building.NameBuilding != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *building.NameBuilding)
		argId++
	}

	if building.NameFullBuilding != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *building.NameFullBuilding)
		argId++
	}

	if building.SelfService != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, *building.SelfService)
		argId++
	}

	if building.AvailabilityAsu != nil {
		setValues = append(setValues, fmt.Sprintf("nameBusinessEntity=$%d", argId))
		args = append(args, *building.AvailabilityAsu)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE buildings SET %s WHERE id=$%d`, setQuery, argId)

	args = append(args, buildingId)

	_, err := r.db.Exec(query, args...)

	return err

}
