create table TypeOfObject
(
    id          serial primary key,
    type_object text
);

create table TradingNetwork
(
    id              serial primary key,
    network_trading text
);

create table Region
(
    id          serial primary key,
    name_region text
);

create table StreetHouse
(
    id           serial primary key,
    street_name  text,
    house_number text
);

create table MicroDistrict
(
    id                  serial primary key,
    micro_district_name text
);

create table FormOfOwnerShip
(
    id        serial primary key,
    form_name text
);


create table BuildingAddress
(
    id                serial primary key,
    region_id         int,
    micro_district_id int,
    street_house_id   int,
    FOREIGN KEY (region_id) REFERENCES Region (id) on delete cascade,
    FOREIGN KEY (micro_district_id) REFERENCES MicroDistrict (id) on delete cascade,
    FOREIGN KEY (street_house_id) REFERENCES StreetHouse (id) on delete cascade
);

create table BuildingEntity
(
    id                  serial primary key,
    name_building       text,
    name_full_building  text,
    object_type         bool default false,
    self_service        bool default false,
    availability_asu    bool default false,
    total_area          int,
    retail_space        int,
    opening_date        date,
    closing_date        date,
    workPlaceCount      int,
    employee_count      int,
    street_name         text,
    micro_district_name text,
    inn                 text,
    kpp                 text,
    region_id           int,
    typeOfObject_id     int,
    tradingNetwork_id   int,
    form_owner_id       int,
    FOREIGN KEY (typeOfObject_id) REFERENCES TypeOfObject (id) on delete cascade,
    FOREIGN KEY (tradingNetwork_id) REFERENCES TradingNetwork (id) on delete cascade,
    FOREIGN KEY (region_id) REFERENCES Region (id) on delete cascade,
    FOREIGN KEY (form_owner_id) REFERENCES FormOfOwnerShip (id) on delete cascade
);

create table Users
(
    id            serial primary key,
    name          text,
    email         text,
    password_hash text
)
