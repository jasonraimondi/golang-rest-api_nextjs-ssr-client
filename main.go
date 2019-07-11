package main

import (
	"fmt"
	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/web"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	ENABLE_DEBUGGING = kingpin.
		Flag("debug", "Enable Debug.").
		Short('D').
		Envar("ENABLE_DEBUGGING").
		Bool()
	JWT_SECURE_KEY = kingpin.
		Flag("jwt-secure-key", "Secure JWT Key, changing this logs everyone out.").
		Short('k').
		Envar("JWT_SECURE_KEY").
		String()
	PG_HOST = kingpin.
		Flag("pg-host", "Postgres Host").
		Short('h').
		Envar("PG_HOST").
		Default("localhost").
		String()
	PG_PORT = kingpin.
		Flag("pg-port", "Postgres Port").
		Short('P').
		Envar("PG_PORT").
		Default("5432").
		String()
	PG_USER = kingpin.
		Flag("pg-user", "Postgres User").
		Short('u').
		Envar("PG_USER").
		Default("print").
		String()
	PG_PASSWORD = kingpin.
		Flag("pg-password", "Postgres Host").
		Short('p').
		Envar("PG_PASSWORD").
		Default("print").
		String()
	PG_DATABASE = kingpin.
		Flag("pg-datbase", "Postgres Host").
		Short('d').
		Envar("PG_DATABASE").
		Default("print").
		String()
)

var (
	a *lib.Application
	h *web.Handler
)

func init() {
	kingpin.Parse()
	s := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*PG_HOST, *PG_PORT, *PG_USER, *PG_PASSWORD, *PG_DATABASE,
	)
	fmt.Println(s)
	dbx, err := sqlx.Connect("postgres", s)
	if err != nil {
		panic(err)
	}
	a = lib.NewApplication(dbx)
	h = web.NewHandler(a, *JWT_SECURE_KEY)
}

func main() {
	e := echo.New()
	e.Debug = *ENABLE_DEBUGGING

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := middleware.JWTConfig{
		Claims:     &web.JwtCustomClaims{},
		SigningKey: []byte(*JWT_SECURE_KEY),
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
