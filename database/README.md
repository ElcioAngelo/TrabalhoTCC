# Estrutura do Banco de Dados 

Este banco de dados foi construído utilizando o processo de **normalização de dados**, com o objetivo de reduzir redundâncias, aumentar a integridade dos dados e melhorar o desempenho geral do sistema.


## Tecnologias Utilizadas

- **PostgreSQL** — Banco de dados relacional
- **PL/pgSQL** — Linguagem procedural para criação de funções e triggers


## Estrutura Geral

### Tipos Especiais (`ENUM`)

São definidos para padronizar valores em colunas específicas:

- `user_role_enum`: `'customer'`, `'admin'`
- `status_enum`: `'active'`, `'inactive'`
- `order_enum`:` 'em processamento'`, `'cancelada'`, `'finalizada'`


## Usuários

#### `users`

Armazena dados básicos do usuário.

| Coluna            | Tipo        | Descrição                        |
|------------------|-------------|----------------------------------|
| id               | serial      | Chave primária                  |
| name             | text        | Nome do usuário                 |
| email            | text        | E-mail (único)                  |
| password         | text        | Senha do usuário                |
| cellphone_number | varchar(15) | Número de telefone              |

---
#### `user_information`

Informações adicionais sobre os usuários.

| Coluna        | Tipo             | Descrição                                   |
|---------------|------------------|---------------------------------------------|
| id            | serial           | Chave primária                             |
| user_id       | int              | Chave estrangeira de usuários (id)`                           |
| creation_date | date             | Data de criação                            |
| update_date   | date             | Data de atualização                        |
| user_role     | user_role_enum   | Cargo: `customer` ou `admin`               |
| status        | status_enum      | Status: `active` ou `inactive`             |

---

#### `user_address_information`

Endereços dos usuários.

| Coluna         | Tipo         | Descrição              |
|----------------|--------------|------------------------|
| id             | serial       | Chave primária        |
| user_id        | int          | Chave estrangeira de usuários (id)|
| state          | text         | Estado do usuário     |
| city           | text         | Cidade do usuário     |
| postal_code    | text         | Código postal         |
| address        | text         | Endereço              |
| address_number | varchar(20)  | Número do endereço    |


## Produtos e Categorias

#### `categories`

Define as categorias dos produtos.

| Coluna | Tipo          | Descrição        |
|--------|---------------|------------------|
| id     | serial        | Chave primária   |
| name   | varchar(100)  | Nome da categoria|

---

#### `products`

Armazena produtos do sistema.

| Coluna      | Tipo         | Descrição                        |
|-------------|--------------|----------------------------------|
| id          | serial       | Chave primária                  |
| name        | varchar(100) | Nome do produto                 |
| price       | float        | Preço                           |
| description | text         | Descrição do produto            |
| status      | status_enum  | Status do produto (`active`, `inactive`) |
| category_id | int          | FK → `categories(id)`           |


## Pedidos

#### `orders`

Armazena os pedidos realizados pelos usuários.

| Coluna     | Tipo         | Descrição                                 |
|------------|--------------|-------------------------------------------|
| id         | serial       | Chave primária                           |
| order_date | date         | Data do pedido                           |
| status     | order_enum   | Status do pedido                         |
| user_id    | int          | FK → `users(id)`                         |

---

### `order_products`

Tabela de relacionamento **muitos-para-muitos** entre pedidos e produtos.

| Coluna     | Tipo   | Descrição                             |
|------------|--------|---------------------------------------|
| id         | serial | Chave primária                       |
| order_id   | int    | FK → `orders(id)`                    |
| product_id | int    | FK → `products(id)`                  |
| quantity   | int    | Quantidade do produto no pedido      |

**Restrição única:** Um mesmo produto só pode aparecer uma vez por pedido (`UNIQUE(order_id, product_id)`).


## Triggers e Funções

### `insert_user_update_date()`

Função **disparada após atualização na tabela `users`**, que atualiza automaticamente o campo `update_date` na tabela `user_information`.

### Trigger: `after_user_update`

- É removida e recriada sempre que o script é executado.
- **Acionada após** qualquer atualização na tabela `users`.


## 🚀 Indexes

Indexes foram criados para melhorar o desempenho de buscas e operações com **chaves estrangeiras**:

```sql
orders(user_id)
products(category_id)
order_products(order_id)
order_products(product_id)
