package handlers

import "github.com/wertick01/dclib/internals/app/models"

func (handler *UsersHandler) CreateUserResponseHelper(u *models.User) *models.UserAfter {
	return &models.UserAfter{
		UserId:     u.UserId,
		Name:       u.Name,
		Surname:    u.Surname,
		Patrynomic: u.Patrynomic,
		Phone:      u.Phone,
		Mail:       u.Mail,
		Role:       u.Role,
	}
}
