package models

type User struct {
	ID        *int64 `json:"id"`
	Name      string `json:"name"`
	SecretKey string `json:"secret_key"`
}
