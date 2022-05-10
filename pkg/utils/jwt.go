package utils

import (
	"errors"
	"time"

	"github.com/amchicas/go-auth-srv/internal/domain"
	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}
type jwtClaims struct {
	Id       int64
	Username string
	Role     int64
	jwt.StandardClaims
}

func (j *JwtWrapper) Sign(user *domain.Auth) (signedToken string, err error) {

	claims := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err

	}
	return signedToken, nil
}
func (j *JwtWrapper) Validate(signedToken string) (claims *jwtClaims, err error) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {

			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Unauthorized")
		}
		return nil, errors.New("Bad request")
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {

		return nil, errors.New("Error Token")

	}
	if claims.ExpiresAt < time.Now().Local().Unix() {

		return nil, errors.New("JWT is expired")
	}
	return claims, nil
}
