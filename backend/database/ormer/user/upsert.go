package ormer

import "main/database/models"

func (uo *userOrmer) Upsert(user models.User) error {

	query := uo.db.Model(&user)

	if user.ID != nil {
		query.OnConflict("(id) DO UPDATE").
			Set("name = EXCLUDED.name").
			Set("secret_key = EXCLUDED.secret_key")
	}

	_, err := query.Insert()
	if err != nil {
		return err
	}

	return nil
}
