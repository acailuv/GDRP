package ormer

import "main/database/models"

func (uo *userOrmer) SelectAll() ([]models.User, error) {
	return []models.User{}, nil
}
