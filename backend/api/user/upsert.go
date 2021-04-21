package api

import (
	"main/database/models"
	"main/utils"
	"net/http"
)

func (u *user) Upsert(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := utils.ReadBody(r.Body, &user)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	err = u.userOrmer.Upsert(user)
	if err != nil {
		utils.InternalServerError(w, "Upsert", err)
		return
	}

	utils.OK(w, "Upsert")
}
