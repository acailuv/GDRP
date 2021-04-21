package api

import (
	ormer "main/database/ormer/user"
	"net/http"

	"github.com/go-pg/pg/v10"
)

type User interface {
	Select(w http.ResponseWriter, r *http.Request)
	SelectAll(w http.ResponseWriter, r *http.Request)
	Upsert(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type user struct {
	userOrmer ormer.UserOrmer
}

func NewUserClient(db *pg.DB) User {
	return &user{
		userOrmer: ormer.NewUserOrmer(db),
	}
}
