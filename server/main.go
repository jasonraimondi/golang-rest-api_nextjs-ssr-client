package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
	"git.jasonraimondi.com/jason/jasontest/app/lib/awsupload"
	"git.jasonraimondi.com/jason/jasontest/app/lib/service"
	"git.jasonraimondi.com/jason/jasontest/app/models"
	"git.jasonraimondi.com/jason/jasontest/server/handlers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

var tables = []interface{}{
	&models.Photo{},
	&models.Tag{},
	&models.User{},
	&models.SignUpConfirmation{},
}

// initialize over init because kingpin.Parse() was causing issues running tests WITH coverage when in the init function
func init() {
	if env("ENABLE_DEBUGGING", "true") == "true" {
		enableDebugging = true
	}
	jwtSecureKey = env("JWT_SECURE_KEY", "my-secret-key")
	dbDriver = env("DB_DRIVER", "postgres")
	dbConnection = env("DB_CONNECTION", "host=localhost port=5432 user=print password=print dbname=print sslmode=disable")
	s3Host = env("S3_HOST", "http://localhost:9000")
	s3Region = env("S3_REGION", "us-east-1")
	s3IdentifierKey = env("S3_IDENTIFIER_KEY", "miniominiominio")
	s3SecretKey = env("S3_SECRET_KEY", "miniominiominio")

	db, err := gorm.Open(dbDriver, dbConnection)
	if err != nil {
		panic("failed to connect to database")
	}
	if enableDebugging {
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))
		migrate(db)
	}
	sessionToken := "" // @todo what is session token?
	s3Config := awsupload.NewS3Config("originals", &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s3IdentifierKey, s3SecretKey, sessionToken),
		Endpoint:         aws.String(s3Host),
		Region:           aws.String(s3Region),
		S3ForcePathStyle: aws.Bool(true),
	})
	app = lib.NewApplication(db, s3Config, jwtSecureKey, dbMigrationsDir)
	h = handlers.NewHandler(app)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(tables...)
	db.Model(&models.Photo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.PhotoTag{}).AddForeignKey("photo_id", "photos(id)", "CASCADE", "CASCADE")
	db.Model(&models.PhotoTag{}).AddForeignKey("tag_id", "tags(id)", "CASCADE", "CASCADE")
	db.Model(&models.SignUpConfirmation{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
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

	e.POST("/login", h.AuthHandler().Login)
	e.POST("/sign_up", h.SignUpHandler().SignUp)
	e.GET("/sign_up_confirmation", h.SignUpHandler().SignUpConfirmation)
	e.GET("/photos/user/:userId", h.Photo().ListForUser)
	e.GET("/photos/tags", h.Photo().ListForTags)
	e.GET("/photos/:photoId", h.Photo().Show)

	admin := e.Group("/admin")
	//admin.Use(authRoute)
	admin.POST("/photos/user/:userId", h.AdminPhoto().Create)
	admin.POST("/photos/:photoId/tags", h.AdminPhoto().AttachTags)
	admin.POST("/photos/:photoId/tags/:tagId", h.AdminPhoto().RemoveTag)

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
