package main

import (
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
				Envar("ENABLE_DEBUGGING").
				Bool()
	JWT_SECURE_KEY = kingpin.
			Flag("jwt-secure-key", "Secure JWT Key, changing this logs everyone out.").
			Envar("JWT_SECURE_KEY").
			String()
	DB_DRIVER = kingpin.
		Flag("db-driver", "Database Driver").
		Envar("DB_DRIVER").
		Default("postgres").
		Enum("postgres", "sqlite3")
	DB_CONNECTION = kingpin.
		Flag("db-connection", "DB Connection").
		Envar("DB_CONNECTION").
		Default("host=localhost port=5432 user=print password=print dbname=print sslmode=disable").
		String()
)

var (
	a *lib.Application
	h *web.Handler
)

func main() {
	initialize()

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
	e.GET("/confirm-email", h.ConfirmEmail)
	e.GET("/", h.Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(authRoute)
	r.GET("", h.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

// initialize over init because kingpin.Parse() was causing issues running tests WITH coverage when in the init function
func initialize() {
	kingpin.Parse()
	dbx, err := sqlx.Connect(*DB_DRIVER, *DB_CONNECTION)
	if err != nil {
		panic(err)
	}
	a = lib.NewApplication(dbx)
	h = web.NewHandler(a, *JWT_SECURE_KEY)
}
