package ormer

import (
	"main/database/models"

	"github.com/go-pg/pg/v10"
)

func (uo *userOrmer) Select(ID int64) (models.User, error) {

	user := models.User{}

	err := uo.db.Model(&user).
		Where("id = ?", ID).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return models.User{}, err
	}

	return user, nil
}
