package auth_service

import (
	"shop/internal/models"
)

type Auth struct {
	Id int64
	Ak string

	DeptId  int64
	Enabled int

	PageNum  int
	PageSize int

	M *models.Auth

	Ids []int64
}

func (a *Auth) DetailByKey(key string) (*models.Auth, error) {
	return models.GetBusinessByKey(key)
}
