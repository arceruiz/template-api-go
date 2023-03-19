package rest

import (
	"net/http"

	"github.com/arceruiz/template-api-go/internal/channels"
	"github.com/arceruiz/template-api-go/internal/channels/rest/middlewares"
	"github.com/arceruiz/template-api-go/internal/config"
	"github.com/arceruiz/template-api-go/internal/integration/dummyIntegration"
	"github.com/arceruiz/template-api-go/internal/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type rest struct {
	dummyService service.DummyService
}

func New() channels.Channel {
	return &rest{
		dummyService: dummyIntegration.New(),
	}
}

func (r *rest) Start() error {
	e := echo.New()
	e.Use(middlewares.NewLogger())
	e.Use(middleware.Recover())

	//e.GET("/jwks", r.JWKS)
	e.GET("/", r.DummyAction)

	return e.Start(":" + config.Instance.Server.HttpPort)
}

func (r *rest) DummyAction(ec echo.Context) error {
	authResponse, err := r.dummyService.DummyAction(nil)
	if err != nil {
		return err //ec.JSON(errWithCode(err))
	}
	return ec.JSON(http.StatusOK, authResponse)
}
