create table users (
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
create table user_information ( 
	id serial primary key,
	user_id int,
	creation_date date default null,
	update_date date default null,
	user_role user_role_enum default 'customer',
	status user_status_enum default 'active',
	foreign key (user_id) references Users(id)
);

create type ratingEnum as enum('1','2','3','4','5');

create table reviews (
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

create table product_categories (
	id serial primary key,
	product_id int,
	category_id int,
	foreign key (product_id) references Products(id),
	foreign key (category_id) references categories(id)
);

create table brands ( 
	id serial primary key,
	name varchar(100)
);

create type "product_status" as enum('active','inactive');

create table products (
	id serial primary key,
	"name" varchar(100) not null,
	price float not null,
	description text not null,
	status product_status,
	category_id int,
	brand_id int,
	foreign key (brand_id) references brands(id)
);

alter table Products
add constraint fk_category_id foreign key 
(category_id) references product_categories(id);


create table item_order (
	id serial primary key,
	product_id int,
	user_id int,
	foreign key (product_id) references Products(id),
	foreign key (user_id) references Users(id)
);

create type orderStatus as enum('pending','processing','delivered','cancelled');
create type paymentMethod as enum('boleto','cartão','dinheiro','pix');

create table orders ( 
	id serial primary key,
	order_date date default null,
	product_quantity int default 0,
	status orderStatus,
	payment_method paymentMethod,
	item_order_id Int,
	foreign key (item_order_id) references Item_order(id)
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

create type operation_enum as enum('exit','entry');

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
