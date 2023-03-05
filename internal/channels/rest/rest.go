package rest

import (
	"github.com/arceruiz/template-api-go/internal/channels"
	"github.com/arceruiz/template-api-go/internal/channels/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type rest struct {
}

func New() channels.Channel {
	return &rest{}
}

func (r *rest) Start() error {
	e := echo.New()
	e.Use(middlewares.NewLogger())
	e.Use(middleware.Recover())

	//e.GET("/jwks", r.JWKS)
	//e.POST("/", r.Authenticate)

	return e.Start(":8081")
}
