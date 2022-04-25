package db

import "golang.org/x/crypto/bcrypt"

//EncryptPassword is the routine that allow to encrypt the password
func EncryptPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}
