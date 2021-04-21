package connection

import (
	"github.com/go-pg/pg/v10"
)

func NewConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "root",
		Password: "root",
		Database: "app-db",
	})
}
