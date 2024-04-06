# GOBANK

<!--toc:start-->

- [GOBANK](#gobank)
  - [tech and tools](#tech-and-tools)
  - [migrate](#migrate) - [up](#up) - [down](#down)
  <!--toc:end-->

## tech and tools

- golang
- database
  - postgresql
- SQL Query Builders
  - [slqc](https://github.com/sqlc-dev/sqlc)
- Database Schema Migration
  - [migrate](https://github.com/golang-migrate/migrate)

## migrate

```bash
 migrate create -ext sql -dir db/migration -seq init_schema
```

### up

```bash
❯ migrate -path db/migration -database "postgresql://<user>:<user_password>@localhost:5432/<db_name>?sslmode=disable" --verbose up
```

### down

```bash
❯ migrate -path db/migration -database "postgresql://<user>:<user_password>@localhost:5432/<db_name>?sslmode=disable" --verbose down
```

## sqlc

### init

```bash
sqlc init
```
