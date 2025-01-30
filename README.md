# CRUD SIMPLE

## Suba o PostgreSQL (exemplo com Docker):
```bash
  docker run --name my-postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=dbname -p 5432:5432 -d postgres
```

## Código SQL
```sql
CREATE TABLE users (
   id SERIAL PRIMARY KEY,          -- ID único, auto-incrementado
   name VARCHAR(100) NOT NULL,     -- Nome do usuário, obrigatório
   email VARCHAR(150) NOT NULL,    -- E-mail do usuário, obrigatório
   created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Data/hora de criação, padrão para horário atual
);
```

## Teste as rotas:
### Criar usuário:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "John", "email": "john@example.com"}' http://localhost:8080/users
Listar usuários:
```
### Listar usuários:
```bash
curl http://localhost:8080/users
```

## Estrutura Final do Projeto
```text
go-crud-postgres/
├── database/
│   └── database.go
├── models/
│   └── models.go
├── repository/
│   └── repository.go
├── handlers/
│   └── handlers.go
├── go.mod
├── go.sum
└── main.go
```

