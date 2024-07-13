package security

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
)

type TokenPair struct {
	AccessToken  string `json:"access_token,omitempty" form:"access_token,omitempty" query:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty" form:"refresh_token,omitempty" query:"refresh_token,omitempty"`
}

type TokenSub struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func JWTFactory() func(*TokenSub) (*TokenPair, *errors.SerivceError) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	return func(sub *TokenSub) (*TokenPair, *errors.SerivceError) {
		var err error

		expiration_short := time.Now().Add(time.Hour * 72)
		expiration_long := time.Now().Add(time.Hour * 24 * 30)

		jwtClaims := jwt.MapClaims{
			"sub": sub,
			"exp": expiration_short.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
		jwt := &TokenPair{}

		jwt.AccessToken, err = token.SignedString([]byte(jwtSecret))
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		}

		return createRefreshToken(jwt, jwtSecret, expiration_long)
	}
}

func ValidateRefreshToken(pair *TokenPair) (*TokenSub, *errors.SerivceError) {
	token, err := jwt.Parse(pair.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	user := TokenSub{}
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Error: err}
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Error: err, Message: "invalid token"}
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err = parser.ParseUnverified(payload["token"].(string), claims)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	payload, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Error: err, Message: "invalid token"}
	}

	if err = mapToStruct(payload["sub"], &user); err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	return &user, nil
}

func ValidateToken(accessToken string) (*TokenSub, *errors.SerivceError) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	user := TokenSub{}
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Error: err}
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if err = mapToStruct(payload["sub"], &user); err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		}

		return &user, nil
	}

	return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Error: err, Message: "invalid token"}
}

func createRefreshToken(token *TokenPair, jwtSecret string, expiration time.Time) (*TokenPair, *errors.SerivceError) {
	var err error

	claims := jwt.MapClaims{}
	claims["token"] = token.AccessToken
	claims["exp"] = expiration.Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.RefreshToken, err = refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	return token, nil
}

func mapToStruct[T comparable](mapped any, structure *T) error {
	str, err := json.Marshal(mapped)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(str, structure); err != nil {
		return err
	}
	return nil
}
