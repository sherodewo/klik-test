package main

import (
	"encoding/gob"
	"fmt"
	"github.com/gorilla/context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"klik/config"
	"klik/config/credential"
	middlewareFunc "klik/middleware"
	"klik/models"
	"klik/routes"
	"klik/utils/session"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("conf/config.env")
	if err != nil {
		log.Fatal("ERROR ", err)
	}

	if err := credential.CredentialsConfig(); err != nil {
		log.Fatal("ERROR ", err)
	}

	gob.Register(session.UserInfo{})
	gob.Register(session.FlashMessage{})
	gob.Register(models.User{})
	gob.Register(models.Menu{})
	gob.Register(map[string]interface{}{})
	gob.Register([]models.ValidationError{})

	//New instance echo
	e := echo.New()

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "auth/error.html", nil)
	}

	e.Static("/klik/assets", "assets")

	e.Pre(middleware.RemoveTrailingSlash())

	//Database
	db := config.NewDbMssql()

	//Validation
	e.Validator = config.NewValidator()

	//Set Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: io.MultiWriter(os.Stdout),
	}))

	e.Use(echo.WrapMiddleware(context.ClearHandler))

	session.Manager = session.NewSessionManager(middlewareFunc.NewCookieStore())

	routes.BackendRoute(e, db)
	routes.FrontendRoute(e, db)
	routes.ApiRoute(e, db)

	// Start server
	if err := e.Start(fmt.Sprintf("%s:%s", credential.AppHost, credential.AppPort)); err != nil {
		e.Logger.Info("shutting down the server")
	}
}
