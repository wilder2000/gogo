package http

import "golang.org/x/crypto/bcrypt"

func EPassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
}
