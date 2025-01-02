package service

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashingPas(pass string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(pass),12)

	if err != nil {
		return "",err
	}

	return string(hash),nil
}

func ParsedTime(t string) (time.Time, error) {
	parsedTime, err := time.Parse("15:04:05", t)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func ParsedDate(t string) (time.Time, error) {
	parsedTime, err := time.Parse(time.DateOnly, t)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
	
}