package db

import "golang.org/x/crypto/bcrypt"

var bcryptRounds = 10

func ComparePassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (m *UsersStorage) hashPassword(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcryptRounds)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
