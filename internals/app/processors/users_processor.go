package processors

import (
	"errors"

	"github.com/wertick01/dclib/internals/app/db"
	"github.com/wertick01/dclib/internals/app/models"
)

type UsersProcessor struct {
	storage *db.UsersStorage
}

func NewUsersProcessor(storage *db.UsersStorage) *UsersProcessor {
	processor := new(UsersProcessor)
	processor.storage = storage
	return processor
}

func (processor *UsersProcessor) CreateUser(user *models.User) (*models.User, error) {

	if user.Name == "" {
		return nil, errors.New("name should not be empty")
	}

	return processor.storage.CreateNewUser(user)
}

func (processor *UsersProcessor) FindUser(id int64) (*models.User, error) {
	user, err := processor.storage.GetUserById(id)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil

}

func (processor *UsersProcessor) FindByPhone(phone string) (*models.User, error) {
	user, err := processor.storage.GetUserByPhone(phone)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil

}

func (processor *UsersProcessor) ListUsers() ([]*models.User, error) {
	return processor.storage.GetUsersList()
}

func (processor *UsersProcessor) UpdateUser(user *models.User) (*models.User, error) { //!!! ПРОВЕРИТЬ

	changeduser, err := processor.storage.ChangeUser(user)
	if err != nil {
		return user, errors.New("SOMETHING IS WRONG")
	}

	return changeduser, nil
}

func (processor *UsersProcessor) DeleteUser(id int64) (*models.User, error) {
	user, _ := processor.FindUser(id)
	_, err := processor.storage.DeleteUserById(id)
	if err != nil {
		return nil, errors.New("CANNOT DELETE USER")
	}
	return user, nil
}
