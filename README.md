# todo-golang
Todo application in golang

## Dependencies
- Web framework - `Gin`
- Database - `PostgreSQL`
- `Docker` for running database
- ORM - `GORM`
- API - `REST API`

## Installation of dependencies
- Database setup
```bash
docker run --name postgres \
  -e POSTGRES_USER=todo \
  -e POSTGRES_PASSWORD=pass \
  -e POSTGRES_DB=todo_db \
  -p 5432:5432 \
  -d postgres
```
- Install go dependencies
```bash
go mod tidy
```

### Endpoints
```
GET     /todos
POST    /todos
PUT     /todos/:id
DELETE  /todos/:id
```



