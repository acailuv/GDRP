package ormer

import "main/database/models"

func (uo *userOrmer) Select(ID int64) (models.User, error) {

	user := models.User{}

	err := uo.db.Model(&user).
		Where("id = ?", ID).
		Select()
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
