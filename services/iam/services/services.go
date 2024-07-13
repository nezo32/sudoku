package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/security"
)

type ServiceContext struct {
	Database       *pgxpool.Pool
	Context        context.Context
	JWTGenerator   func(*security.TokenSub) (*security.TokenPair, *errors.SerivceError)
	PasswordEncode func(string) (string, error)
	Echo           *echo.Echo
}

type UsersOutput struct {
	ID        uuid.UUID  `sql:"primary_key" json:"id,omitempty" xml:"id"`
	Username  string     `json:"username,omitempty" xml:"username"`
	FirstName string     `json:"first_name,omitempty" xml:"first_name"`
	LastName  string     `json:"last_name,omitempty" xml:"last_name"`
	Email     string     `json:"email,omitempty" xml:"email"`
	CreatedAt time.Time  `json:"created_at,omitempty" xml:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" xml:"updated_at"`
}
