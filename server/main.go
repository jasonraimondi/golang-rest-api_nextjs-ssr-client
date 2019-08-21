package main

import (
	"net/http"
	"os"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
	"git.jasonraimondi.com/jason/jasontest/app/lib/awsupload"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
	"git.jasonraimondi.com/jason/jasontest/server/handlers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	enableDebugging bool
	jwtSecureKey    string
	dbDriver        string
	dbConnection    string
	dbMigrationsDir string
	s3Host          string
	s3Region        string
	s3IdentifierKey string
	s3SecretKey     string
	app             *lib.Application
	h               *handlers.Handler
)

// initialize over init because kingpin.Parse() was causing issues running tests WITH coverage when in the init function
func init() {
	if env("ENABLE_DEBUGGING", "true") == "true" {
		enableDebugging = true
	}
	jwtSecureKey = env("JWT_SECURE_KEY", "my-secret-key")
	dbDriver = env("DB_DRIVER", "postgres")
	dbConnection = env("DB_CONNECTION", "host=localhost port=5432 user=print password=print dbname=print sslmode=disable")
	dbMigrationsDir = env("DB_MIGRATIONS_DIR", "/Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations")
	s3Host = env("S3_HOST", "http://localhost:9000")
	s3Region = env("S3_REGION", "us-east-1")
	s3IdentifierKey = env("S3_IDENTIFIER_KEY", "miniominiominio")
	s3SecretKey = env("S3_SECRET_KEY", "miniominiominio")

	dbx, err := sqlx.Connect(dbDriver, dbConnection)
	if err != nil {
		panic(err)
	}
	sessionToken := "" // @todo what is session token?
	s3Config := awsupload.NewS3Config("originals", &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3IdentifierKey, s3SecretKey, sessionToken),
		Endpoint:         aws.String(s3Host),
		Region:           aws.String(s3Region),
		S3ForcePathStyle: aws.Bool(true),
	})
	app = lib.NewApplication(dbx, s3Config, jwtSecureKey, dbMigrationsDir)
	h = handlers.NewHandler(app)
}

func main() {
	e := echo.New()
	e.Debug = enableDebugging

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	}))

	config := middleware.JWTConfig{
		Claims:       &service.JwtCustomClaims{},
		SigningKey:   []byte(jwtSecureKey),
		ErrorHandler: func(err error) error { return err },
	}
	authRoute := middleware.JWTWithConfig(config)

	e.POST("/login", h.Auth().Login)
	e.POST("/sign_up", h.SignUp().SignUp)
	e.GET("/sign_up_confirmation", h.SignUp().SignUpConfirmation)
	e.GET("/list_photos", h.Photo().List)

	guard := e.Group("/api")
	guard.Use(authRoute)
	guard.POST("/upload", h.Photo().Upload)

	e.Logger.Fatal(e.Start(":1323"))
}

func env(env string, fallback string) (result string) {
	result = os.Getenv(env)
	if result == "" {
		result = fallback
	}
	return result
}
