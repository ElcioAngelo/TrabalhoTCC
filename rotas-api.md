## Orders

### GET /sales (admin) 
Obs: orders com status SOLD groupby MONTH (isso é para o dashboard gráficos)

### GET /revenue (admin) 
Obs: Valores dos pedidos do tipo VENDA com status SOLD - Valores dos pedidos do tipo COMPRA com status BUYED 

### GET /orders (admin)

### POST /sales (user)


## Stock

### POST /stock (admin) Exemplo: Tipo, venda, perda, roubo, compra

### GET /stock (admin) 
Exemplo listagem de todas as movimentações de estoque (Tudo o que entrou - tudo o que saiu)

## Users

### GET /user (admin)
### GET /users (admin)
### POST /user (admin)
### PUT /user (admin)

## Auth

### POST /auth 

#### Request {user, password}
#### Response
{
    user: {
        ...dados do usuário
    },
    token: tokenJwt
}


## Cadastros

user:~$ sudo -i -u postgres
postgres@user:~$ psql

ALTER USER postgres PASSWORD 'senha_aqui';