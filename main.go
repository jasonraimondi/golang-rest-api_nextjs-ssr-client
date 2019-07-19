package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"git.jasonraimondi.com/jason/jasontest/handlers"
	"git.jasonraimondi.com/jason/jasontest/lib"
	"git.jasonraimondi.com/jason/jasontest/lib/s3"
)

var (
	EnableDebugging bool
	JwtSecureKey string
	DbDriver string
	DbConnection string
	DbMigrationsDir string
	S3Host string
	S3Region string
	S3IdentifierKey string
	S3SecretKey string
	a *lib.Application
	h *handlers.Handler
)

// initialize over init because kingpin.Parse() was causing issues running tests WITH coverage when in the init function
func init() {
	if env("ENABLE_DEBUGGING", "true") == "true" {
		EnableDebugging = true
	}
	JwtSecureKey = env("JWT_SECURE_KEY", "my-secret-key")
	DbDriver = env("DB_DRIVER", "postgres")
	DbConnection = env("DB_CONNECTION", "host=localhost port=5432 user=print password=print dbname=print sslmode=disable")
	DbMigrationsDir = env("DB_MIGRATIONS_DIR", "/Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations")
	S3Host = env("S3_HOST", "http://localhost:9000")
	S3Region = env("S3_REGION", "us-east-1")
	S3IdentifierKey = env("S3_IDENTIFIER_KEY", "miniominiominio")
	S3SecretKey = env("S3_SECRET_KEY", "miniominiominio")

	dbx, err := sqlx.Connect(DbDriver, DbConnection)
	if err != nil {
		panic(err)
	}
	sessionToken := "" // @todo what is session token?
	s3Config := s3.NewS3Config("originals", &aws.Config{
		Credentials:      credentials.NewStaticCredentials(S3IdentifierKey, S3SecretKey, sessionToken),
		Endpoint:         aws.String(S3Host),
		Region:           aws.String(S3Region),
		S3ForcePathStyle: aws.Bool(true),
	})
	a = lib.NewApplication(dbx, s3Config, JwtSecureKey, DbMigrationsDir)
	h = handlers.NewHandler(a)
}

func env(env string, fallback string) (result string) {
	result = os.Getenv(env)
	if result == "" {
		result = fallback
	}
	return result
}

func main() {
	e := echo.New()
	e.Debug = EnableDebugging

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := middleware.JWTConfig{
		Claims:     &handlers.JwtCustomClaims{},
		SigningKey: []byte(JwtSecureKey),
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
