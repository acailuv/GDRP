package ormer

import "main/database/models"

func (uo *userOrmer) Select(ID int64) (models.User, error) {
	return models.User{}, nil
}
