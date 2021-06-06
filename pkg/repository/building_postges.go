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

	fmt.Println("Micro: " + building.MicroDistrictName)
	fmt.Println("Street: " + building.StreetName)
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

func (r *BuildingPostgres) GetAll(nameBuilding, typeOfObject, networkTrading, region string) ([]domain.BuildingSelect, error) {
	var items []domain.BuildingSelect

	//SELECT
	//       bt.id, bt.name_building, bt.object_type, bt.self_service, bt.availability_asu, bt.total_area,
	//       bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count,
	//       tf.object_type, tn.network_trading, rg.name_region, bt.street_name, bt.micro_district_name
	//       FROM BuildingEntity bt, TypeOfObject tf, TradingNetwork tn, Region rg
	//WHERE bt.region_id = rg.id and bt.typeOfObject_id = tf.id and bt.tradingNetwork_id = tn.id;

	if nameBuilding == ""  && typeOfObject == "0" && networkTrading == "0" && region == "0" {
		fmt.Println("Without filtering")
		query := fmt.Sprintf(`SELECT bt.id, bt.name_building, bt.name_full_building, bt.object_type, bt.self_service, bt.availability_asu,
       bt.total_area, bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count, bt.street_name, bt.micro_district_name,
       bt.inn, bt.kpp, tf.type_object, tn.network_trading, rg.name_region, fs.form_name
       from BuildingEntity bt JOIN TypeOfObject tf on bt.typeOfObject_id = tf.id JOIN TradingNetwork tn on bt.tradingNetwork_id = tn.id
          JOIN Region rg on bt.region_id=rg.id JOIN FormOfOwnerShip fs on bt.form_owner_id = fs.id;`)
		if err := r.db.Select(&items, query); err != nil {
			return nil, err
		}
	} else {
		fmt.Println("filtering")

		/*
		bt.street_name, bt.micro_district_name
		SELECT bt.id, bt.name_building, bt.name_full_building, bt.object_type, bt.self_service, bt.availability_asu,
		       bt.total_area, bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count,
		       bt.inn, bt.kpp, tf.type_object, tn.network_trading, rg.name_region, fs.form_name
		       from BuildingEntity bt JOIN TypeOfObject tf on bt.typeOfObject_id = tf.id JOIN TradingNetwork tn on bt.tradingNetwork_id = tn.id
		          JOIN Region rg on bt.region_id=rg.id JOIN FormOfOwnerShip fs on bt.form_owner_id = fs.id where tn.id = 0 or tf.id = 0 or rg.id=3 or bt.name_building='';
		 */
		query := fmt.Sprintf(`SELECT bt.id, bt.name_building, bt.object_type, bt.self_service, bt.availability_asu, bt.total_area,
        bt.retail_space, bt.opening_date, bt.closing_date, bt.workPlaceCount, bt.employee_count, bt.street_name, bt.micro_district_name, bt.inn, bt.kpp, 
        tf.type_object, tn.network_trading, rg.name_region, fs.form_name
		from %s bt JOIN %s tf on bt.typeOfObject_id = tf.id JOIN
    	%s tn on bt.tradingNetwork_id = tn.id JOIN %s rg on bt.region_id = rg.id JOIN %s fs on bt.form_owner_id = fs.id
		WHERE bt.name_building = $1 or rg.id = $2 or tf.id = $3 or
		tn.id = $4`, buildingEntityTable, typeOfObjectsTable, tradingNetworkTable, regionTable, formOfOwnerShipTable)

		//or bt.typeOfObject_id=$2 or bt.tradingNetwork_id=$3 or bt.region_id=$4 or bt.street_name=$5 or bt.micro_district_name > $6

		if err := r.db.Select(&items, query, nameBuilding, region, typeOfObject, networkTrading); err != nil {
			return nil, err
		}
	}

	return items, nil

	//, typeOfObject, networkTrading, region, streetName, microDistrict
}

func (r *BuildingPostgres) Delete(buildingId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, "BuildingEntity")
	_, err := r.db.Exec(query, buildingId)
	return err
}

func (r *BuildingPostgres) Update(buildingId int, building domain.BuildingUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if building.NameBuilding != nil {
		setValues = append(setValues, fmt.Sprintf("name_building=$%d", argId))
		args = append(args, *building.NameBuilding)
		argId++
	}

	if building.NameFullBuilding != nil {
		setValues = append(setValues, fmt.Sprintf("name_full_building=$%d", argId))
		args = append(args, *building.NameFullBuilding)
		argId++
	}

	if building.ObjectType != nil {
		setValues = append(setValues, fmt.Sprintf("object_type=$%d", argId))
		args = append(args, *building.ObjectType)
		argId++
	}

	if building.SelfService != nil {
		setValues = append(setValues, fmt.Sprintf("self_service=$%d", argId))
		args = append(args, *building.SelfService)
		argId++
	}

	if building.AvailabilityAsu != nil {
		setValues = append(setValues, fmt.Sprintf("availability_asu=$%d", argId))
		args = append(args, *building.AvailabilityAsu)
		argId++
	}

	if building.TotalArea != nil {
		setValues = append(setValues, fmt.Sprintf("total_area=$%d", argId))
		args = append(args, *building.TotalArea)
		argId++
	}

	if building.RetailSpace != nil {
		setValues = append(setValues, fmt.Sprintf("retail_space=$%d", argId))
		args = append(args, *building.RetailSpace)
		argId++
	}

	if building.OpeningDate != nil {
		setValues = append(setValues, fmt.Sprintf("opening_date=$%d", argId))
		args = append(args, *building.OpeningDate)
		argId++
	}

	if building.ClosingDate != nil {
		setValues = append(setValues, fmt.Sprintf("closing_date=$%d", argId))
		args = append(args, *building.ClosingDate)
		argId++
	}

	if building.WorkPlaceCount != nil {
		setValues = append(setValues, fmt.Sprintf("workPlaceCount=$%d", argId))
		args = append(args, *building.WorkPlaceCount)
		argId++
	}

	if building.EmployeeCount != nil {
		setValues = append(setValues, fmt.Sprintf("employee_count=$%d", argId))
		args = append(args, *building.EmployeeCount)
		argId++
	}

	if building.StreetName != nil {
		setValues = append(setValues, fmt.Sprintf("street_name=$%d", argId))
		args = append(args, *building.StreetName)
		argId++
	}

	if building.MicroDistrictName != nil {
		setValues = append(setValues, fmt.Sprintf("micro_district_name=$%d", argId))
		args = append(args, *building.MicroDistrictName)
		argId++
	}

	if building.Inn != nil {
		setValues = append(setValues, fmt.Sprintf("inn=$%d", argId))
		args = append(args, *building.Inn)
		argId++
	}

	if building.Kpp != nil {
		setValues = append(setValues, fmt.Sprintf("kpp=$%d", argId))
		args = append(args, *building.Kpp)
		argId++
	}

	if building.RegionId != nil {
		setValues = append(setValues, fmt.Sprintf("region_id=$%d", argId))
		args = append(args, *building.RegionId)
		argId++
	}

	if building.TypeObjectId != nil {
		setValues = append(setValues, fmt.Sprintf("typeOfObject_id=$%d", argId))
		args = append(args, *building.TypeObjectId)
		argId++
	}

	if building.TradingNetworkId != nil {
		setValues = append(setValues, fmt.Sprintf("tradingNetwork_id=$%d", argId))
		args = append(args, *building.TradingNetworkId)
		argId++
	}

	if building.FormOwnerId != nil {
		setValues = append(setValues, fmt.Sprintf("form_owner_id=$%d", argId))
		args = append(args, *building.FormOwnerId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE BuildingEntity SET %s WHERE id=$%d`, setQuery, argId)

	args = append(args, buildingId)

	_, err := r.db.Exec(query, args...)

	return err

}
