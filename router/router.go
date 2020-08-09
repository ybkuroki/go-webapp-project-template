package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-project-template/config"
	"github.com/ybkuroki/go-webapp-project-template/controller"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo, conf *config.Config) {
	if conf.Extension.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
			},
			MaxAge: 86400,
		}))
	}

	e.HTTPErrorHandler = controller.JSONErrorHandler
	e.Use(middleware.Recover())

	e.GET(controller.APIAccountLoginStatus, controller.GetLoginStatus())
	e.GET(controller.APIAccountLoginAccount, controller.GetLoginAccount())

	if conf.Extension.SecurityEnabled {
		e.POST(controller.APIAccountLogin, controller.PostLogin())
		e.POST(controller.APIAccountLogout, controller.PostLogout())
	}

	e.GET(controller.APIHealth, controller.GetHealthCheck())

}
