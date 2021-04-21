package api

import (
	"main/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (u *user) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.InternalServerError(w, "Parse Int", err)
		return
	}

	err = u.userOrmer.Delete(ID)
	if err != nil {
		utils.InternalServerError(w, "Delete", err)
		return
	}

	utils.OK(w, "Delete")
}
