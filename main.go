package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"

	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/web"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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
	S3_ID = kingpin.
		Flag("s3-id", "DB Connection").
		Envar("S3_ID").
		Default("miniominiominio").
		String()
	S3_SECRET = kingpin.
		Flag("s3-secret", "DB Connection").
		Envar("S3_SECRET").
		Default("miniominiominio").
		String()
)

var (
	a *lib.Application
	h *web.Handler
)

// initialize over init because kingpin.Parse() was causing issues running tests WITH coverage when in the init function
func initialize() {
	kingpin.Parse()
	dbx, err := sqlx.Connect(*DB_DRIVER, *DB_CONNECTION)
	if err != nil {
		panic(err)
	}
	sessionToken := ""
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(*S3_ID, *S3_SECRET, sessionToken),
		//Endpoint:         aws.String("https://s3.wasabisys.com"),
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
	}
	a = lib.NewApplication(dbx, s3Config)
	h = web.NewHandler(a, *JWT_SECURE_KEY)
}

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
	e.POST("/upload", h.Upload)

	//It is just like express javascript
	e.POST("/sign-up", h.SignUp)
	e.GET("/confirm-email", h.ConfirmEmail)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(authRoute)
	r.GET("", h.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
