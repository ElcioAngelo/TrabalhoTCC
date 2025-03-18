# Trabalho TCC Back End

Para melhor visualização do trabalho recomendo que instale essas extensões do VSCODE:

* Better Comments (Aaron Bond)
* Pacote de extensão "Golang" (aldijav) 

#### Atualização 14/03:

* Configuração do banco de dados (conexão)

* Implementação do usecase (Equivalente ao Services)

* Implemetação do repository

* Implemetanção do controller

* Modelos Completos
(11 modelos)

* pasta CMD (arquivo principal)

* Implementando arquivos Docker (Banco de dados, Aplicativo principal)

#### Atualização 16/03:

* Proteção contra SQL injections.
* Senhas são criptografadas (Usando o Bcrypt)
* Deixando comentários para o código ficar mais legivel (main, product_controller, product_repo,user_repo)

#### Atualização 17/03:

* Requisições get,post,delete criado para produtos e usuarios.




## Implementações futuras:

* Implementar Patch para ambos usuários e produtos
* Implementar Autenticação JWT
* adcionar funções para administradores
* terminar Patch do usuário.

## Banco de dados:

#### Users Table
| Column             | Data Type         | Constraints                                |
|--------------------|-------------------|--------------------------------------------|
| `id`               | `serial`          | `PRIMARY KEY`                              |
| `name`             | `varchar(100)`     | `NOT NULL`                                 |
| `email`            | `varchar(100)`     | `UNIQUE`, `NOT NULL`                       |
| `password`         | `varchar(100)`     | `NOT NULL`                                 |
| `cellphone_number` | `varchar(15)`      | `NOT NULL`                                 |
| `shipping_adress`  | `varchar(100)`     | `UNIQUE`, `NOT NULL`                       |
| `payment_address`  | `varchar(100)`     | `UNIQUE`, `NOT NULL`                       |

#### User_information Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `user_id`     | `int`             | `FOREIGN KEY (user_id) REFERENCES Users(id)` |
| `creation_date` | `date`           | `DEFAULT NULL`                             |
| `update_date` | `date`            | `DEFAULT NULL`                             |
| `user_role`   | `user_role_enum`  | `DEFAULT 'customer'`                       |
| `status`      | `user_status_enum`| `DEFAULT 'active'`                         |

#### Reviews Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `user_id`     | `int`             | `FOREIGN KEY (user_id) REFERENCES Users(id)` |
| `description` | `text`            |                                            |
| `rating`      | `ratingEnum`      |                                            |

#### Categories Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `name`        | `varchar(100)`     |                                            |

#### Brands Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `name`        | `varchar(100)`     |                                            |

#### Products Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `name`        | `varchar(100)`     | `NOT NULL`                                 |
| `price`       | `float`           | `NOT NULL`                                 |
| `description` | `text`            | `NOT NULL`                                 |
| `category_id` | `int`             | `FOREIGN KEY (category_id) REFERENCES categories(id)` |
| `brand_id`    | `int`             | `FOREIGN KEY (brand_id) REFERENCES brands(id)` |

#### Item_order Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `product_id`  | `int`             | `FOREIGN KEY (product_id) REFERENCES Products(id)` |
| `user_id`     | `int`             | `FOREIGN KEY (user_id) REFERENCES Users(id)` |

#### Orders Table
| Column            | Data Type         | Constraints                                |
|-------------------|-------------------|--------------------------------------------|
| `id`              | `serial`          | `PRIMARY KEY`                              |
| `order_date`      | `date`            | `DEFAULT NULL`                             |
| `product_quantity`| `int`             | `DEFAULT 0`                                |
| `status`          | `orderStatus`     |                                            |
| `payment_method`  | `paymentMethod`   |                                            |
| `item_order_id`   | `int`             | `FOREIGN KEY (item_order_id) REFERENCES Item_order(id)` |

#### Favorites Table
| Column        | Data Type         | Constraints                                |
|---------------|-------------------|--------------------------------------------|
| `id`          | `serial`          | `PRIMARY KEY`                              |
| `product_id`  | `int`             | `FOREIGN KEY (product_id) REFERENCES Products(id)` |
| `user_id`     | `int`             | `FOREIGN KEY (user_id) REFERENCES Users(id)` |

#### Batch Table
| Column            | Data Type         | Constraints                                |
|-------------------|-------------------|--------------------------------------------|
| `id`              | `serial`          | `PRIMARY KEY`                              |
| `code`            | `varchar(100)`     |                                            |
| `expiration_date` | `date`            | `DEFAULT NULL`                             |

#### Stock_moviment Table
| Column            | Data Type         | Constraints                                |
|-------------------|-------------------|--------------------------------------------|
| `id`              | `serial`          | `PRIMARY KEY`                              |
| `moviment_date`   | `date`            | `DEFAULT NULL`                             |
| `product_id`      | `int`             | `FOREIGN KEY (product_id) REFERENCES Products(id)` |
| `quantity`        | `int`             |                                            |
| `batch_id`        | `int`             | `FOREIGN KEY (batch_id) REFERENCES batch(id)` |
| `value`           | `float`           |                                            |
| `operation`       | `operationEnum`   |                                            |

---

#### Types

#### `user_role_enum`
| Value     |
|-----------|
| `guest`   |
| `admin`   |
| `customer`|

#### `user_status_enum`
| Value     |
|-----------|
| `active`  |
| `inactive`|
| `suspended`|

#### `ratingEnum`
| Value |
|-------|
| `1`   |
| `2`   |
| `3`   |
| `4`   |
| `5`   |

#### `orderStatus`
| Value      |
|------------|
| `pending`  |
| `processing`|
| `delivered` |
| `cancelled` |

#### `paymentMethod`
| Value     |
|-----------|
| `boleto`  |
| `cartão`  |
| `dinheiro`|
| `pix`     |

#### `operationEnum`
| Value     |
|-----------|
| `exit`    |
| `entry`   |


### Código SQL

```
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

create table Products (
	id serial primary key,
	"name" varchar(100) not null,
	price float not null,
	description text not null,
	category_id int,
	brand_id int,
	foreign key (category_id) references categories(id),
	foreign key (brand_id) references brands(id)
);

create table Item_order (
	id serial primary key,
	product_id int,
	user_id int,
	foreign key (product_id) references Products(id),
	foreign key (user_id) references Users(id)
);

create type orderStatus as enum('pending','processing','delivered','cancelled');
create type paymentMethod as enum('boleto','cartão','dinheiro','pix');

create table Orders ( 
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

```

## Instalação do projeto:

```
    #clone o repositorio.
    git clone https://github.com/ElcioAngelo/TrabalhoTCC.git
    #mudando para a pasta salva do respositorio
    cd TrabalhoTCC
    #mudando para a pasta de execucao principal
    cd cmd 
    #rodando o servidor
    go run main.go
```

## Tecnologias usadas

* Golang 
* Docker
* Postgresql

