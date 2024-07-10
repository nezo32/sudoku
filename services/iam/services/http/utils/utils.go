package utils

import (
	"path"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

type EndpointEntry struct {
	base_url string
	Echo     *echo.Echo
	Database *pg.DB
}

type EndpointHandler func(ctx echo.Context, entry *EndpointEntry) error

func CreateEndpointFactory(base_url string, echo *echo.Echo, database *pg.DB) *EndpointEntry {
	return &EndpointEntry{
		base_url,
		echo,
		database,
	}
}

func (entry *EndpointEntry) GET(url string, handler EndpointHandler) {
	entry.Echo.GET(path.Join(entry.base_url, url), func(c echo.Context) error { return handler(c, entry) })
}

func (entry *EndpointEntry) POST(url string, handler EndpointHandler) {
	entry.Echo.POST(path.Join(entry.base_url, url), func(c echo.Context) error { return handler(c, entry) })
}

func (entry *EndpointEntry) PUT(url string, handler EndpointHandler) {
	entry.Echo.PUT(path.Join(entry.base_url, url), func(c echo.Context) error { return handler(c, entry) })
}

func (entry *EndpointEntry) DELETE(url string, handler EndpointHandler) {
	entry.Echo.DELETE(path.Join(entry.base_url, url), func(c echo.Context) error { return handler(c, entry) })
}
