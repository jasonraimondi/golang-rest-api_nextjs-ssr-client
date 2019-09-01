package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
	"git.jasonraimondi.com/jason/jasontest/app/lib/config"
	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
	"git.jasonraimondi.com/jason/jasontest/server/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	debug         bool
	jwtSecureKey  config.JWTSecureKey
	dbCredentials config.DBCred
	s3Cred        config.S3Cred
)

// initialize over init because kingpin.Parse() was causing issues running tests WITH coverage when in the init function
func init() {
	if env("ENABLE_DEBUGGING", "true") == "true" {
		debug = true
	}
	jwtSecureKey = config.JWTSecureKey(env("JWT_SECURE_KEY", "my-secret-key"))
	dbCredentials = config.DBCred{
		Driver:     env("DB_DRIVER", "postgres"),
		Connection: env("DB_CONNECTION", "host=localhost port=5432 user=print password=print dbname=print sslmode=disable"),
	}
	s3Cred = config.S3Cred{
		Host:       env("S3_HOST", "http://localhost:9000"),
		Region:     env("S3_REGION", "us-east-1"),
		Identifier: env("S3_IDENTIFIER_KEY", "miniominiominio"),
		Secret:     env("S3_SECRET_KEY", "miniominiominio"),
	}
}

func main() {
	db, err := gorm.Open(dbCredentials.Driver, dbCredentials.Connection)
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	var s3Config = config.NewS3Config("originals", &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3Cred.Identifier, s3Cred.Secret, s3Cred.SessionToken),
		Endpoint:         aws.String(s3Cred.Host),
		Region:           aws.String(s3Cred.Region),
		S3ForcePathStyle: aws.Bool(true),
	})
	var app = lib.NewApplication(db, s3Config, jwtSecureKey, debug)
	var h = handlers.NewHandler(app)

	var e = echo.New()
	e.Debug = debug

	if debug {
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))
		repository.Migrate(db)
	}

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

	e.POST("/login", h.AuthHandler().Login)
	e.POST("/sign_up", h.SignUpHandler().SignUp)
	e.GET("/sign_up_confirmation", h.SignUpHandler().SignUpConfirmation)
	e.GET("/photos/user/:userId", h.Photo().ListForUser)
	e.GET("/photos/tags", h.Photo().ListForTags)
	e.GET("/photos/:photoId", h.Photo().Show)

	admin := e.Group("/admin")
	//admin.Use(authRoute)
	admin.POST("/photos/user/:userId", h.AdminPhoto().Create)
	admin.POST("/photos/:photoId/tags/:tagId", h.AdminPhoto().RemoveTag)

	admin.POST("/photos/:photoId", h.AdminPhoto().UpdatePhoto)

	// @todo remove this
	fake := e.Group("/fake")
	fake.Use(authRoute)

	e.Logger.Fatal(e.Start(":1323"))
}

func env(env string, fallback string) (result string) {
	result = os.Getenv(env)
	if result == "" {
		result = fallback
	}
	return result
}
