package api

import (
	"main/utils"
	"net/http"
)

func (u *user) SelectAll(w http.ResponseWriter, r *http.Request) {

	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	users, err := u.userOrmer.SelectAll()
	if err != nil {
		utils.InternalServerError(w, "Select All", err)
		return
	}

	utils.OKWithData(w, users)

}
