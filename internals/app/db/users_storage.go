package db

import (
	"errors"

	"github.com/wertick01/dclib/internals/app/models"

	"database/sql"
)

type UsersStorage struct {
	DB *sql.DB
}

func NewUsersStorage(db *sql.DB) *UsersStorage {
	storage := new(UsersStorage)
	storage.DB = db
	return storage
}

func (m *UsersStorage) CreateNewUser(user *models.User) (*models.User, error) {

	stmt := `INSERT INTO dclib_test.users (username, usersurname, userpatrynomic, userphone, useremail, userhash, userrole) VALUES(?, ?, ?, ?, ?, ?, ?)`

	hash, err := m.hashPassword(user.Hash)
	if err != nil {
		return nil, err
	}
	user.Hash = hash

	result, err := m.DB.Exec(
		stmt,
		user.Name,
		user.Surname,
		user.Patrynomic,
		user.Phone,
		user.Mail,
		user.Hash,
		user.Role.RoleId,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.UserId = id

	return user, nil
}

func (m *UsersStorage) GetUsersList() ([]*models.User, error) {

	stmt := `SELECT u.userid, u.username, u.usersurname, u.userpatrynomic, u.userphone, u.useremail, u.userrole, r.user_role FROM dclib_test.users AS u RIGHT JOIN dclib_test.roles AS r ON u.userrole = r.role_id WHERE u.userid IS NOT NULL`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		s := &models.User{}
		err = rows.Scan(&s.UserId, &s.Name, &s.Surname, &s.Patrynomic, &s.Phone, &s.Mail, &s.Role.RoleId, &s.Role.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UsersStorage) GetUserById(id int64) (*models.User, error) {

	stmt := `SELECT u.userid, u.username, u.usersurname, u.userpatrynomic, u.userphone, u.useremail, u.userhash, u.userrole, r.USER_role FROM dclib_test.users AS u RIGHT JOIN dclib_test.roles AS r ON u.userrole = r.role_id WHERE u.userid = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.User{}

	err := row.Scan(&s.UserId, &s.Name, &s.Surname, &s.Patrynomic, &s.Phone, &s.Mail, &s.Hash, &s.Role.RoleId, &s.Role.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *UsersStorage) GetUserByPhone(phone string) (*models.User, error) {

	stmt := `SELECT u.userid, u.username, u.usersurname, u.userpatrynomic, u.userphone, u.useremail, u.userhash, u.userrole, r.user_role FROM dclib_test.users AS u RIGHT JOIN dclib_test.roles AS r ON u.userrole = r.role_id WHERE u.userphone = ?`

	row := m.DB.QueryRow(stmt, phone)

	s := &models.User{}

	err := row.Scan(&s.UserId, &s.Name, &s.Surname, &s.Patrynomic, &s.Phone, &s.Mail, &s.Hash, &s.Role.RoleId, &s.Role.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *UsersStorage) ChangeUser(old *models.User) (*models.User, error) {

	stmt := `UPDATE dclib_test.users SET username = ?, usersurname = ?, userpatrynomic = ?, userphone = ?, useremail = ?, userhash = ?, userrole = ? WHERE userid = ?`

	hash, err := m.hashPassword(old.Hash)
	if err != nil {
		return nil, err
	}
	old.Hash = hash

	change, err := m.DB.Exec(stmt, old.Name, old.Surname, old.Patrynomic, old.Phone, old.Mail, old.Hash, old.Role.RoleId, old.UserId)
	if err != nil {
		return nil, err
	}

	_, err = change.LastInsertId()
	if err != nil {
		return nil, err
	}

	return old, nil
}

func (m *UsersStorage) DeleteUserById(id int64) (int, error) {
	stmt := `DELETE FROM dclib_test.users WHERE userid = ?`
	deleted, err := m.DB.Exec(stmt, id)
	if err != nil {
		return 0, err
	}

	res, err := deleted.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(res), nil
}
