# GOBANK

<!--toc:start-->

- [GOBANK](#gobank)
  - [tech and tools](#tech-and-tools)
  - [migrate](#migrate) - [up](#up) - [down](#down)
  <!--toc:end-->

## tech and tools

- golang
- postgresql
- migrate [https://github.com/golang-migrate/migrate]

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
