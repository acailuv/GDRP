# GDRP: Go-Docker-React-Postgres
A Template Stack. Use this if you want to spend less time doing infra stuff and more time to actually developing stuff.

# Setting Up
- Install [Go](https://golang.org/doc/install).
- Install [NodeJS](https://nodejs.org/en/download/).
- Install [Docker](https://docs.docker.com/get-docker/).
- Install [Docker Compose](https://docs.docker.com/compose/install/).
- Install [golang-migrate](https://github.com/golang-migrate/migrate).
    - You (might) need [Brew](https://brew.sh/).

# Cheat Sheet
This section contains some commands and examples that you might need.

## Useful Commands
Below are some useful commands that you might need.

### Export POSTGRESQL_URL Environment Variable
```bash
export POSTGRESQL_URL='postgres://root:root@localhost:5432/app-db?sslmode=disable'
```

### Create Migration
```bash
migrate create -ext sql -dir ./backend/database/migrations -seq create_users_table
```

### Run "Up" Migration
```bash
migrate -database ${POSTGRESQL_URL} -path ./backend/database/migrations up
```

### Run "Down" Migration
```bash
migrate -database ${POSTGRESQL_URL} -path ./backend/database/migrations down
```

## Useful Examples
Below are some useful examples that you might need.

### Database Transactions (SQL)
What is a transaction? Find out [here](https://specialties.bayt.com/en/specialties/q/222277/what-is-the-difference-between-transaction-and-query/).
```sql
BEGIN;

CREATE TYPE enum_mood AS ENUM (
	'happy',
	'sad',
	'neutral'
);
ALTER TABLE users ADD COLUMN mood enum_mood;

COMMIT;
```
```sql
BEGIN;

ALTER TABLE users DROP COLUMN mood;
DROP TYPE enum_mood;

COMMIT;
```

### Run Migration From Go File(s)
```golang
import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:postgres@localhost:5432/example?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
```

## Useful Tools
Below are some useful tools that you can use to develop your app.

### Adminer
This is a database admin panel (similar to phpmyadmin). To access:
- Open up your browser.
- `http://localhost:3645`. Or you can click [here](http://localhost:3645).

### ReqBin
This is a free API testing tools to test out your endpoints. To access:
- Install [ReqBin Extension for Chrome](https://chrome.google.com/webstore/detail/reqbin-http-client/gmmkjpcadciiokjpikmkkmapphbmdjok).
- Go to [ReqBin](https://reqbin.com/).