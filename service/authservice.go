package service

import (
	"time"

	"github.com/Zyprush18/resto-go/model/entity"
	"github.com/Zyprush18/resto-go/repositories/databases"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email string
	Password string
}

type TokenPayload struct {
	Email string
	jwt.RegisteredClaims	
}


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

// login function use jwt

func LoginService(email, password string) (*TokenPayload, error) {
	user := new (entity.User)
	if err := databases.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	payload := &TokenPayload{Email: email}

	return payload, nil
}

func CreateToken(payload *TokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil	
}