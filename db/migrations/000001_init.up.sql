create table users (
    id serial unique,
    login varchar(255) unique not null,
    hash_pass varchar(255) not null,
    username varchar(255) not null
);

create table items_type (
    title varchar(255) not null unique
);

insert into items_type values('printer','monitor','storage');

create table items (
    id serial unique,
    item_id integer not null,
    items_type varchar(255) not null references items_type(title),
    in_stock boolean not null
);

create table storeges_type(
    title varchar(55)
);

insert into storages_type values('HDD','SSD');

create table storages (
    id serial unique,
    title varchar(255) not null,
    volume integer not null,
    size varchar(255) not null,
    type varchar(55) not null
);