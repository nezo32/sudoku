package utils

import (
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/services"
)

type EndpointEntry struct {
	Group *echo.Group
	*services.ServiceContext
}

type EndpointHandler func(ctx echo.Context, entry *EndpointEntry) error

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func CreateEndpointFactory(ctx *services.ServiceContext, group *echo.Group) *EndpointEntry {
	return &EndpointEntry{
		group,
		ctx,
	}
}

func (entry *EndpointEntry) GET(url string, handler EndpointHandler, middlewares ...echo.MiddlewareFunc) {
	entry.Group.GET(url, func(c echo.Context) error { return handler(c, entry) }, middlewares...)
}

func (entry *EndpointEntry) POST(url string, handler EndpointHandler, middlewares ...echo.MiddlewareFunc) {
	entry.Group.POST(url, func(c echo.Context) error { return handler(c, entry) }, middlewares...)
}

func (entry *EndpointEntry) PUT(url string, handler EndpointHandler, middlewares ...echo.MiddlewareFunc) {
	entry.Group.PUT(url, func(c echo.Context) error { return handler(c, entry) }, middlewares...)
}

func (entry *EndpointEntry) DELETE(url string, handler EndpointHandler, middlewares ...echo.MiddlewareFunc) {
	entry.Group.DELETE(url, func(c echo.Context) error { return handler(c, entry) }, middlewares...)
}
