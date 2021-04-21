package ormer

import "main/database/models"

func (uo *userOrmer) Delete(ID int64) error {

	user := models.User{}

	_, err := uo.db.Model(&user).
		Where("id = ?", ID).
		Delete()
	if err != nil {
		return err
	}

	return nil
}
