# Trabalho TCC Back End

Para uma melhor visualização do trabalho, recomendo que instale as seguintes extensões do VSCode:

- Better Comments (Aaron Bond)  
- Pacote de extensão "Golang" (aldijav)

## Estrutura do trabalho

- **CMD** – contém o arquivo de inicialização do servidor  
- **Model** – contém a representação das tabelas do banco de dados  
- **Python** – servidor de imagens  
- **DB** – conexão com o banco de dados  
- **Middleware** – função de autenticação utilizando cookies  
- **Repository** – camada responsável pela execução de código SQL  
- **Controller** – camada de requisições HTTP

## Atualização 14/03:

- Configuração do banco de dados  
- Implementação do use case  
- Implementação do repository  
- Implementação do controller  
- Modelos completos (11 modelos)  
- Pasta CMD  

## Atualização 16/03:

- Senhas são criptografadas (usando o Bcrypt)

## Atualização 17/03:

- Requisições GET, POST e DELETE criadas para produtos e usuários

## Atualização 24/03:

- Implementação de rotas para alteração de produto (editar nome, categoria etc.)

## Atualização 29/03:

- Implementação de rotas para usuários:
  - Edição de valores do usuário em rotas separadas. 
- Implementação de rotas para produtos:
  - Edição de valores do produto em rotas separadas.  
- Produtos e usuários não são removidos do sistema; em vez disso, seu status é alterado para 'ativo' ou 'inativo'.
- Produtos e usuários com status inativo não aparecem para usuários comuns.


## Atualização 10/06

* 


## Implementações futuras:
* Implementar um chatbot (inteligencia artificial) para ajudar duvidas dos clientes 


## Instalação do projeto:

```
    git clone https://github.com/ElcioAngelo/TrabalhoTCC.git
    cd TrabalhoTCC
    cd cmd 
    go run main.go
```

## Tecnologias usadas

* Golang 
* Python
* Postgresql
* Docker 

## Autor
* Elcio Angelo Negri
* Email: elcio.negri@grupointegrado.br


