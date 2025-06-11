create table Users (
	id serial primary key,
	"name" varchar(100) not null,
	email varchar(100) unique not null,
	"password" varchar(100) not null,
	cellphone_number varchar(15) not null,
	shipping_adress varchar(100) unique not null,
	payment_address varchar(100) unique not null
);

create type user_role_enum as enum('guest','admin','customer');
create type "user_status_enum" as enum('active','inactive','suspended'); 

/* !! Dividi a tabela do usuário para se tornar mais facil !! */
create table User_information ( 
	id serial primary key,
	user_id int,
	creation_date date default null,
	update_date date default null,
	user_role user_role_enum default 'customer',
	status user_status_enum default 'active',
	foreign key (user_id) references Users(id)
);

CREATE OR REPLACE FUNCTION insert_user_information()
RETURNS TRIGGER AS $$
BEGIN
    -- Insert a new record into User_information
    INSERT INTO User_information (user_id, creation_date, update_date, user_role, status)
    VALUES (NEW.id, CURRENT_DATE, NULL, 'customer', 'active');
    
    -- Return the new row (standard practice for INSERT triggers)
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

create or replace function insert_user_update_date()
returns trigger as $$
begin
	insert into User_information (update_date)
	values (CURRENT_DATE);

	return new;
end;
$$ language plpgsql;

create or replace function 

create trigger after_user_update
after update on Users
for each row 
execute function insert_user_update_date();


CREATE TRIGGER after_user_insert
AFTER INSERT ON Users
FOR EACH ROW
EXECUTE FUNCTION insert_user_information();


create type ratingEnum as enum('1','2','3','4','5');

create table Reviews (
	id serial primary key,
	user_id int,
	description text,
	rating ratingEnum,
	foreign key (user_id) references Users(id)
);

/*########## FIM DAS TABELAS DO USUÁRIO ####################################################################################################*/


/*########## TABELAS RELACIONADAS AO PRODUTO ####################################################################################################*/

create table categories (
	id serial primary key,
	name varchar(100)
);

create table brands ( 
	id serial primary key,
	name varchar(100)
);

create type "product_status" as enum('active','inactive');

create table Products (
	id serial primary key,
	"name" varchar(100) not null,
	price float not null,
	description text not null,
	status product_status,
	brand_id int,
	category_id int,
	foreign key (category_id) references categories(id),
	foreign key (brand_id) references brands(id)
);

create table user_order (
	id serial primary key,
	user_id int,
	quantity int,
	status varchar(10),
	foreign key (user_id) references Users(id),
);

create type orderStatus as enum('pending','processing','delivered','cancelled');
create type orderStatus2 as enum('pending','cancelled','done');
create type paymentMethod as enum('boleto','cartão','dinheiro','pix');

create table Orders ( 
	id serial primary key,
	order_date timestamp default null,
	payment_method paymentMethod,
	status orderStatus2,
	user_id int,
	foreign key (user_id) references users(id)
);

create table favorites ( 
	id serial primary key,
	product_id int,
	user_id int,
	foreign key (product_id) references Products(id),
	foreign key (user_id) references Users(id)
);

/*########## FIM DAS TABELAS DO PRODUTO ####################################################################################################*/

/*########## TABELAS DO ESTOQUE  ####################################################################################################*/

create table batch (
	id serial primary key,
	code varchar(100),
	expiration_date date default null
);

create type operationEnum as enum('exit','entry');

create table stock_moviment (
	id serial primary key,
	moviment_date date default null,
	product_id int,
	quantity int,
	batch_id int,
	value float,
	operation operationEnum,
	foreign key (product_id) references products(id),
	foreign key (batch_id) references batch(id)
);

/*########## FIM DAS TABELAS DO ESTOQUE  ####################################################################################################*/

insert into Brands ("name")
values('adidas'),('nike'),('marca teste'),('marca teste2');

insert into categories ("name") 
values('Computers'),('Mobile Hardware'),('Music Equipment');

select * from categories;
select * from products;

insert into product_categories (product_id,category_id)
values (18,1),(19,2),(20,3);

INSERT INTO Products ("name", price, description, status,brand_id)
VALUES
  ('Laptop', 1500.00, 'High-performance laptop', 'active', 1),
  ('Smartphone', 800.00, 'Latest model smartphone', 'active', 2),
  ('Headphones', 150.00, 'Wireless noise-cancelling headphones', 'active', 2);

-- Inactive products
INSERT INTO Products ("name", price, description, status, brand_id)
VALUES
  ('Old Smartphone', 200.00, 'Outdated smartphone', 'inactive', 1),
  ('Old Laptop', 300.00, 'Old laptop model', 'inactive', 1);

select * from products;

select * from product_categories;

select p.id, 
p."name", 
p.description,p.price,
c."name" as "product_category", 
br."name" as "product_brand",
p.status
from Products p
inner join product_categories prc on p.id = prc.product_id
inner join brands br on br.id = p.brand_id
inner join categories c on c.id = prc.category_id;

select * from products;

select p.id, p."name", p.description, p.price, c."name" as "product_category", br."name" as "product_brand", p.status as "product_status"
		from Products p
		inner join product_categories prc on p.id = prc.product_id
		inner join brands br on br.id = p.brand_id
		inner join categories c on c.id = prc.category_id;
	

select * from product_categories;

insert into product_categories (product_id,category_id)
values(1,2),(2,3),(4,1),(5,1),(6,2),(9,1);


select * from orders;

alter table orders
drop column product_quantity;

alter table orders 
drop column item_order_id;

alter table item_order
drop column user_id;

alter table item_order
add column quantity int;

alter table item_order
add column price int;

alter table orders
add column user_id int;

alter table orders
add constraint user_oder_fk foreign key
(user_id) references users(id);

alter table item_order
add constraint item_order_fk foreign key 
(order_id) references orders(id);

select * from item_order;


create or replace function insert_user_update_date()
returns trigger as $$
begin
	insert into User_information (update_date)
	values (CURRENT_DATE);

	return new;
end;
$$ language plpgsql;


select * from orders;

insert into orders (order_date,status,payment_method,user_id)
values(current_date,)


/* SCRIPT PARA TESTES */
create table sales (
	id serial primary key,
	user_id int,
	sale_date timestamp,
	foreign key (user_id) references users(id)
);

select * from sales;

create or replace function insert_sale()
returns trigger as $$
begin
	/* a utilização de "old.status is distinct" é para evitar que vendas sejam duplicadas */
	if new.status = 'done' and old.status is distinct from 'done' then
		insert into sales (user_id,sale_date)
		values (new.user_id, current_timestamp);
	end if;
	
	return new;
end;
$$ language plpgsql;

create or replace trigger inserting_done_orders
after update on orders
for each row 
execute function insert_sale();

INSERT INTO brands (name) VALUES
  ('Quimiclor'),
  ('AquaSan'),
  ('PoolClean'),
  ('ChloroMax'),
  ('EcoPools');

INSERT INTO categories (name) VALUES
  ('Desinfetantes'),
  ('Controladores'),
  ('Especiais'),
  ('Algicidas');

INSERT INTO products (name, price, description, status, category_id) VALUES
  ('Desinfetante Cloro Pro 10L', 79.90, 'Desinfetante concentrado com alto teor de cloro ativo.', 'active',1),
  ('Controlador de pH Balanceador 5L', 45.50, 'Produto para estabilizar o pH da água da piscina.', 'active',2),
  ('Algicida de Choque Turbo 2L', 65.00, 'Remove rapidamente algas verdes e amarelas.', 'active',4),
  ('Desinfetante Multiuso Premium 1L', 12.90, 'Ideal para limpeza de superfícies e sanitização geral.', 'active',1),
  ('Controle de Cloro Automático 20kg', 229.00, 'Granulado de liberação lenta para manutenção de cloro.', 'active',2),
  ('Desinfetante de Uso Geral Eco 5L', 35.99, 'Fórmula biodegradável e eficiente.', 'active',1),
  ('Algicida Manutenção Sem Espuma 1L', 29.90, 'Previne formação de algas sem formar espuma.', 'active',4),
  ('Produto Especial Brilho Água Cristal 1L', 39.90, 'Deixa a água da piscina cristalina e reluzente.', 'active',3),
  ('Controlador de Alcalinidade Plus 3L', 37.50, 'Corrige a alcalinidade total da água.', 'active',2),
  ('Desinfetante Clorado Extra Forte 20L', 140.00, 'Alta concentração para desinfecção pesada.', 'active',1);


select * from products;
-- Product categories
INSERT INTO  (product_id, category_id) VALUES
  (1, 1), (4, 1), (6, 1), (10, 1),     -- Desinfetantes
  (9, 2), (5, 2), (2, 2),             -- Controladores
  (3, 4), (7, 4),                     -- Algicidas
  (8, 3);                             -- Especiais
  
select p.id, p.name, p.price, p.description, p.status, c.name as "product_category" 
		from products p
		join categories c on c.id = p.category_id
where lower(c."name") = lower('Desinfetantes');

 select p.id, p."name", p.description, p.price, c."name" as "product_category", 
               p.status as "product_status"
        from Products p
        inner join categories c on c.id = p.category_id;

  select * from products;
  
  select * from users
 
  select s.user_id as "user_id",  u."name" as "name",
  u.email as "email",
  u.cellphone_number as "user_cellphone",
  u.payment_address as "user_payment_address",
  u.shipping_adress as "user_shipping_address",
  s.sale_date as "sale_date",
  s.total_revenue as "revenue"
  from sales s 
  inner join users u on s.user_id = u.id;

  select * from users;

