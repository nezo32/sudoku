package utils

import (
	"path"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/services"
)

type EndpointEntry struct {
	base_url string
	*services.ServiceContext
}

type EndpointHandler func(ctx echo.Context, entry *EndpointEntry) error

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func CreateEndpointFactory(base_url string, ctx *services.ServiceContext) *EndpointEntry {
	return &EndpointEntry{
		base_url,
		ctx,
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
