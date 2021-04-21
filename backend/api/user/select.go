package api

import (
	"main/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *user) Select(w http.ResponseWriter, r *http.Request) {

	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.InternalServerError(w, "Parse Int", err)
		return
	}

	user, err := u.userOrmer.Select(ID)
	if err != nil {
		utils.InternalServerError(w, "Select", err)
		return
	}

	utils.OKWithData(w, user)

}
