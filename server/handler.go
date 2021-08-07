package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sharpvik/luba/rrlist"
	"net/http"
)

// Runtime holds all the things we need at... runtime.
type Runtime struct {
	nodes *rrlist.RoundRobinList
}

func NewRuntime() *Runtime {
	return &Runtime{
		nodes: rrlist.New(),
	}
}

func (rt *Runtime) Handler() (e *echo.Echo) {
	e = echo.New()
	e.GET("/", index)
	return
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "index")
}
