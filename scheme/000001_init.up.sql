create table buildings
(
    id serial not null unique,
    name varchar(255) not null,
    address varchar(255) not null,
    phone varchar(15) not null unique,
    name_business_entity varchar(255) not null
)

