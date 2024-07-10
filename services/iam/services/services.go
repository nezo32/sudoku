package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type ServiceContext struct {
	Database *pgxpool.Pool
	Context  context.Context
	Echo     *echo.Echo
}
