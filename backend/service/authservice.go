package service

import (
	"fmt"
	"time"

	"github.com/Zyprush18/resto-go/backend/model/entity"
	"github.com/Zyprush18/resto-go/backend/repositories/databases"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type keyRole string
var RoleKey keyRole = "key_role"
var UserIdKey keyRole = "user_id"

type Login struct {
	Email    string `json:"email"`
	Password string	`json:"password"`
}

type TokenPayload struct {
	Role string
	jwt.RegisteredClaims
}

func HashingPas(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 12)

	if err != nil {
		return "", err
	}
	return string(hash), nil
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
	user := new(entity.User)
	if err := databases.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}


	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	payload := &TokenPayload{
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: fmt.Sprintf("%d",user.ID),
			Subject: user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

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


func ParsedToken(getToken string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(getToken, &TokenPayload{},func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"),nil
	})

	if parsToken,ok := token.Claims.(*TokenPayload); ok {
		return parsToken,nil
	}

	return nil,err
}