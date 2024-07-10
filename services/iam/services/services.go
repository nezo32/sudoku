package services

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

type ServiceContext struct {
	Database *pg.DB
	Echo     *echo.Echo
}
