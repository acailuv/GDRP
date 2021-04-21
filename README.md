# GDRP: Go-Docker-React-Postgres
A Template Stack. Use this if you want to spend less time doing infra stuff and more time to actually developing stuff.

# Setting Up
- Install [Go](https://golang.org/doc/install).
- Install [NodeJS](https://nodejs.org/en/download/).
- Install [Docker](https://docs.docker.com/get-docker/).
- Install [Docker Compose](https://docs.docker.com/compose/install/).
- Install [golang-migrate](https://github.com/golang-migrate/migrate).
    - You (might) need [Brew](https://brew.sh/).

# Run
First time running? Use `docker-compose up --build`.\
If not, use `docker-compose up`.

**NOTE:** You need to rebuild if you made some changes to the Dockerfile(s). However, I'm pretty sure you know this if you are planning to make changes to those file(s).

# Cheat Sheet
This section contains some commands and examples that you might need. I have made a [Makefile](https://github.com/acailuv/GDRP/blob/main/Makefile) to make things easier. Check it out!

## Useful Commands
Below are some useful commands that you might need.

### Using Makefile Macros
```bash
make generate-migration name=AddUserTable
```
```bash
make run-migration-up database=${POSTGRESQL_URL}
```
```bash
make run-migration-down database=${POSTGRESQL_URL}
```

### Export POSTGRESQL_URL Environment Variable
```bash
export POSTGRESQL_URL='postgres://root:root@localhost:5432/app-db?sslmode=disable'
```
Not Working? Try these steps:
- Open `~/.bashrc` using `nano` or your favorite text editor.
- Scroll to the very bottom.
- Copy the command above and paste it there.
- Save and exit.
- `source ~/.bashrc` or log out and log back in.

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

# Errors
In this section I will list out some errors that I encountered and provide solutions for that problem.

## I/O Timeout when trying to `go get` in Docker
The error will look like this:
```
go: github.com/go-pg/pg/v10@v10.9.1: Get "https://proxy.golang.org/github.com/go-pg/pg/v10/@v/v10.9.1.mod": dial tcp: lookup proxy.golang.org on 172.17.254.1:53: read udp 172.17.0.2:45074->172.17.254.1:53: i/o timeout
```
To solve:
- `nano /etc/resolv.conf`
- Change all `nameserver` to 8.8.8.8 and 8.8.4.4
- Your `resolv.conf` should looks like this (more or less)
```
nameserver 8.8.8.8
nameserver 8.8.4.4
```

## CORS Problem when trying to call API
When you call an endpoint from the frontend, this error will appear. The error will look something like this (browser console):
```
Access to fetch at 'http://localhost:5000/users' from origin 'http://localhost:3000' has been blocked by CORS policy: Request header field access-control-allow-origin is not allowed by Access-Control-Allow-Headers in preflight response.
```
To solve:
- Install [Moesif](https://chrome.google.com/webstore/detail/moesif-origin-cors-change/digfbfaphojjndkpccljibejjbppifbc) in Google Chrome.
- Turn the extension on
- Refresh the React page.