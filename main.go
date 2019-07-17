package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"

	"git.jasonraimondi.com/jason/jasontest/domain/lib"
	"git.jasonraimondi.com/jason/jasontest/domain/service"
	"git.jasonraimondi.com/jason/jasontest/web"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	EnableDebugging = kingpin.
		Flag("debug", "Enable Debug.").
		Envar("ENABLE_DEBUGGING").
		Bool()
	JwtSecureKey = kingpin.
		Flag("jwt-secure-key", "Secure JWT Key, changing this logs everyone out.").
		Envar("JWT_SECURE_KEY").
		String()
	DbDriver = kingpin.
		Flag("db-driver", "Database Driver").
		Envar("DB_DRIVER").
		Default("postgres").
		Enum("postgres", "sqlite3")
	DbConnection = kingpin.
		Flag("db-connection", "DB Connection").
		Envar("DB_CONNECTION").
		Default("host=localhost port=5432 user=print password=print dbname=print sslmode=disable").
		String()
	S3Host = kingpin.
		Flag("s3-host", "S3 Origin").
		Envar("S3_HOST").
		Default("http://localhost:9000").
		String()
	S3Region = kingpin.
		Flag("s3-region", "S3 Region").
		Envar("S3_REGION").
		Default("us-east-1").
		Enum("us-east-1")
	S3IdentifierKey = kingpin.
		Flag("s3-id", "S3 Identifier Key").
		Envar("S3_IDENTIFIER_KEY").
		Default("miniominiominio").
		String()
	S3SecretKey = kingpin.
		Flag("s3-secret", "S3 Secret Key").
		Envar("S3_SECRET_KEY").
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
	dbx, err := sqlx.Connect(*DbDriver, *DbConnection)
	if err != nil {
		panic(err)
	}
	sessionToken := "" // @todo what is session token?
	s3Config := service.NewS3Config("originals", &aws.Config{
		Credentials:      credentials.NewStaticCredentials(*S3IdentifierKey, *S3SecretKey, sessionToken),
		Endpoint:         aws.String(*S3Host),
		Region:           aws.String(*S3Region),
		S3ForcePathStyle: aws.Bool(true),
	})
	a = lib.NewApplication(dbx, s3Config)
	h = web.NewHandler(a, *JwtSecureKey)
}

func main() {
	initialize()

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
