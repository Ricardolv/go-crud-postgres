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

# Diagramas de Sequência

## Endpoint: GET /user

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Database

    Client->>API: GET /user
    API->>Database: Query user data
    Database-->>API: User data
    API-->>Client: Response with user data
```

---

## Endpoint: POST /login

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Database

    Client->>API: POST /login (credentials)
    API->>Database: Verify user credentials
    alt Credentials valid
        Database-->>API: User authenticated
        API-->>Client: 200 OK (Session token)
    else Credentials invalid
        Database-->>API: Invalid credentials
        API-->>Client: 401 Unauthorized
    end
```

# Diagrama C4 - Representação em Mermaid

## Contexto - Sistema CRUD com PostgreSQL

```mermaid
C4Context
    Person(Usuario, "Usuário", "Interage com o sistema via API para gerenciar dados de usuários.")
    System(SystemCRUD, "Sistema CRUD", "Fornece funcionalidades para CRUD de usuários.")
    System_Ext(ExternalDatabase, "PostgreSQL", "Banco de dados usado para persistência dos dados.")

    Rel(Usuario, SystemCRUD, "Realiza requisições via API")
    Rel(SystemCRUD, ExternalDatabase, "Armazena e consulta dados de usuários")
```

---

## Container - Estrutura do Sistema CRUD

```mermaid
C4Container
    System_Boundary(SystemCRUD, "Sistema CRUD") {
        Container(WebApp, "Aplicação Web", "Go + Gorilla/Mux", "Expõe endpoints REST para CRUD de usuários.")
        ContainerDb(Database, "PostgreSQL", "Banco de Dados", "Armazena dados de usuários.")
    }
    Person(Usuario, "Usuário", "Interage com o sistema via API para gerenciar dados de usuários.")
    Rel(Usuario, WebApp, "Requisições HTTP")
    Rel(WebApp, Database, "Consultas e operações de CRUD")
```

---

## Componente - Detalhes da Aplicação

```mermaid
C4Component
    Container_Boundary(WebApp, "Aplicação Web") {
        Component(Handler, "Handlers HTTP", "Expõe e gerencia endpoints REST.")
        Component(Repository, "Repositório", "Abstrai operações sobre o banco de dados.")
        Component(DBConnection, "Conexão PostgreSQL", "Gerencia conexões com o banco de dados.")
    }
    ContainerDb(Database, "PostgreSQL", "Banco de Dados", "Armazena dados de usuários.")

    Person(Usuario, "Usuário", "Interage com a aplicação via API.")
    Rel(Usuario, Handler, "Realiza requisições HTTP")
    Rel(Handler, Repository, "Chama métodos do repositório")
    Rel(Repository, DBConnection, "Usa para efetuar operações no banco de dados")
    Rel(DBConnection, Database, "Executa comandos SQL")
```

---

## Código - Fluxo de Controle

```mermaid
sequenceDiagram
    participant Usuario
    participant API
    participant DB

    Usuario->>API: Envia requisição HTTP (ex: GET /users)
    API->>DB: Executa Query (ex: SELECT * FROM users)
    DB-->>API: Retorna dados da consulta
    API-->>Usuario: Retorna resposta HTTP com dados
```

---

### Explicação dos elementos

1. **Diagrama de Contexto**:
    - Mostra como os usuários interagem com o sistema e como este se conecta ao banco PostgreSQL.

2. **Diagrama de Container**:
    - Detalha os principais componentes, destacando a aplicação feita em Go que interage com o banco de dados.

3. **Diagrama de Componente**:
    - Divide a aplicação Go em partes específicas:
        - `Handlers HTTP`: Responsáveis pelos endpoints REST.
        - `Repositórios`: Abstrações para acesso ao banco.
        - `Conexões com o Banco`: Gerenciamento direto das transações no PostgreSQL.

4. **Diagrama de Sequência**:
    - Explica o fluxo de uma requisição, desde o cliente até as interações no banco, e a resposta ao cliente.

Se precisar de algo mais específico ou detalhado, como adicionar mais elementos ou melhorar o contexto para sua aplicação, basta avisar! 😊

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

# go-crud-postgres
