package security

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenSub struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func JWTFactory(secret []byte) func(TokenSub) (*TokenPair, *errors.SerivceError) {
	sec := secret

	return func(claims TokenSub) (*TokenPair, *errors.SerivceError) {

		expiration_short := time.Now().Add(time.Hour * 72)
		expiration_long := time.Now().Add(time.Hour * 24 * 30)

		access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": claims,
			"exp": expiration_short.Unix(),
		})
		refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": claims.ID,
			"exp": expiration_long.Unix(),
		})
		access_signed, err := access_token.SignedString(sec)
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Message: "Can't sign JWT token", Error: err}
		}
		refresh_signed, err := refresh_token.SignedString(sec)
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Message: "Can't sign JWT token", Error: err}
		}

		return &TokenPair{
			AccessToken:  access_signed,
			RefreshToken: refresh_signed,
		}, nil
	}
}
