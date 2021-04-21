package ormer

import (
	"main/database/models"
)

func (uo *userOrmer) Upsert(user models.User) error {

	res, err := uo.db.Model(&user).WherePK().UpdateNotZero()
	if err != nil {
		return err
	}

	if res.RowsAffected() > 0 {
		return nil
	}

	_, err = uo.db.Model(&user).WherePK().Insert()
	if err != nil {
		return err
	}

	return nil
}
