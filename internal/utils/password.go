package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plainPwd string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)

	return string(result), err
}

func CompareHashAndPlainPassword(hashPwd, plainPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd)) == nil
}
