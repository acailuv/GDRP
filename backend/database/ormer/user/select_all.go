package ormer

import "main/database/models"

func (uo *userOrmer) SelectAll() ([]models.User, error) {

	users := make([]models.User, 0)

	err := uo.db.Model(&users).Select()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}
