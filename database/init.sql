/*
    O banco foi construido usando o processo 
    de organização chamado normalização de dados,
    um conjunto de regras  para reduzir 
    a redundância de dados, aumentar a integridade 
    de dados e o desempenho. 
*/

-- * Tipos especiais -- 
create type user_role_enum as enum('customer','admin');
create type status_enum as enum('active','inactive');
create type order_enum as enum('em processamento','cancelada','finalizada');


-- * Tabela de usuários --  
create table if not exists users (
    id serial primary key,
    name text not null,
    email text not null unique,
    password text not null,
    cellphone_number varchar(15) not null
);

-- * Tabela de informações sobre os usuários --
create table if not exists user_information (
    id serial primary key,
    user_id int not null,
    creation_date date not null,
    update_date date default null,
    user_role user_role_enum not null,
    status status_enum not null,
    foreign key (user_id) references users(id)
);

-- * Tabela de informações sobre o endereço dos usuários -- 
create table if not exists user_address_information (
    id serial primary key,
    user_id int not null,
    state text not null,
	city text not null,
	postal_code text not null,
	address text not null,
	address_number varchar(20) not null,
    foreign key (user_id) references users(id)
);


-- ? Função para armazenar a data de atualização
-- ? dos usuários
create or replace function insert_user_update_date()
returns trigger as $$ 
begin
        -- ? Faz uma atualização na data de atualização do usuário
        -- ? onde o usuário for 
        update user_information
        set update_date = CURRENT_DATE
        where user_id = NEW.id;
        return NEW;
end;
$$ language plpgsql;

-- ? Caso exista já exista um disparador, ele é excluido e será criado um novo.
drop trigger if exists after_user_update on users;

-- ? Dispara a função 
create trigger after_user_update
after update on users
for each row 
execute function insert_user_update_date();

-- * tabela de categorias --
create table if not exists categories (
        id serial primary key,
        name varchar(100) not null
);

-- * Tabela de produtos --
create table if not exists products (
    id serial primary key,
    name varchar(100),
    price float not null,
    description text not null,
    status status_enum not null default 'active',
    category_id int not null,
    foreign key (category_id) references categories(id)
);

-- * Tabela de pedidos --
create table if not exists orders (
    id serial primary key,
    order_date date not null default CURRENT_DATE,
    status order_enum not null default 'em processamento',
    user_id int not null,
    foreign key (user_id) references users(id)
);

-- * Tabela de muitos para muitos de pedidos e produtos. --
create table if not exists order_products (
    id serial primary key,
    order_id int not null,
    product_id int not null,
    quantity int not null,
    foreign key (order_id) references orders(id),
    foreign key (product_id) references products(id),
    unique(order_id, product_id)
);

-- ? Criação de indexes para as chaves estrangeiras
-- ? melhorando buscas e ações CRUDs 
create index if not exists idx_orders_user_id ON orders(user_id);
create index if not exists idx_products_category_id ON products(category_id);
create index if not exists idx_order_products_order_id ON order_products(order_id);
create index if not exists idx_order_products_product_id ON order_products(product_id);

