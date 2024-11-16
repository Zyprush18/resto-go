package service

import "golang.org/x/crypto/bcrypt"

func HashingPas(pass string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(pass),12)

	if err != nil {
		return "",err
	}

	return string(hash),nil
}