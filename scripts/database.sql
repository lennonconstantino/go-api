create table product (
	id SERIAL primary key,
	product_name VARCHAR(5) not null,
	price NUMERIC(10,2) not null
);

insert into product(product_name, price) values('Sushi', 100);
