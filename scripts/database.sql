drop table if exists products;
drop table if exists users;

create table products (
	id SERIAL primary key,
	name VARCHAR(64) not null,
	price NUMERIC(10,2) not null
);

CREATE TABLE users (
    id serial primary key,
    username VARCHAR(50) not null,
    email VARCHAR(50) not null,
    password VARCHAR(100) not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Unique(email)
);

insert into products(name, price) values('Sushi', 100);
insert into products(name, price) values('Rice', 5099);
insert into products(name, price) values('Rib', 3300);
insert into products(name, price) values('Steak', 5099);
insert into products(name, price) values('Pizza', 989);

