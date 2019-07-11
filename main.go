package main

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/web"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	EnableDebugging = kingpin.
		Flag("debug", "Enable Debug.").
		Short('d').
		Bool()
	JwtSecureKey = kingpin.
		Flag("jwt-secure-key", "Secure JWT Key, changing this logs everyone out.").
		Short('k').
		String()
	PG_HOST = kingpin.
		Flag("pg-host", "Postgres Host").
		Short('h').
		Default("localhost").
		String()
	PG_PORT = kingpin.
		Flag("pg-host", "Postgres Port").
		Short('P').
		Default("5432").
		String()
	PG_USER = kingpin.
		Flag("pg-host", "Postgres Database").
		Short('u').
		Default("prints").
		String()
	PG_PASSWORD = kingpin.
		Flag("pg-host", "Postgres Host").
		Short('p').
		Default("prints").
		String()
	PG_DATABASE = kingpin.
		Flag("pg-host", "Postgres Host").
		Short('d').
		Default("prints").
		String()
)

var (
	a *lib.Application
	h *web.Handler
)

func init() {
	kingpin.Parse()
	s := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PG_HOST, PG_PORT, PG_USER, PG_PASSWORD, PG_DATABASE,
	)
	dbx, err := sqlx.Connect("postgres", s)
	if err != nil {
		panic(err)
	}
	a = lib.NewApplication(dbx)
	h = &web.Handler{App: a}
}

func main() {
	e := echo.New()
	e.Debug = *EnableDebugging

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := middleware.JWTConfig{
		Claims:     &web.JwtCustomClaims{},
		SigningKey: []byte(*JwtSecureKey),
	}
	authRoute := middleware.JWTWithConfig(config)

	e.POST("/login", h.Login)

	// Unauthenticated route
	e.POST("/sign-up", h.SignUp)
	e.GET("/", h.Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(authRoute)
	r.GET("", h.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
