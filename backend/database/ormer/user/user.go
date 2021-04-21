package ormer

import (
	"main/database/models"

	"github.com/go-pg/pg/v10"
)

type UserOrmer interface {
	Select(ID int64) (models.User, error)
	SelectAll() ([]models.User, error)
	Upsert(user models.User) error
	Delete(ID int64) error
}

type userOrmer struct {
	db *pg.DB
}

func NewUserOrmer(db *pg.DB) UserOrmer {
	return &userOrmer{
		db: db,
	}
}
