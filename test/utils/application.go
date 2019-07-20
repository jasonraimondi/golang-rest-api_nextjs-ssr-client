package utils

import (
	"database/sql"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"git.jasonraimondi.com/jason/jasontest/handlers"
	"git.jasonraimondi.com/jason/jasontest/lib"
	"git.jasonraimondi.com/jason/jasontest/lib/s3"
)

func NewTestApplication() (a *lib.Application) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	dbx := sqlx.NewDb(db, "sqlite3")
	databaseInstance, err := sqlite3.WithInstance(dbx.DB, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}
	sessionToken := ""
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("", "", sessionToken),
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
	}
	s3Config := s3.NewS3Config("test-originals", config)
	a = lib.NewApplication(
		dbx,
		s3Config,
		"jwtSecureKey-test",
		"/Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations",
	)
	if err = MigrateNow(&databaseInstance, a.MigrationDir); err != nil {
		panic(err)
	}
	return a
}


func NewTestHandler() *handlers.Handler {
	a := NewTestApplication()
	return &handlers.Handler{
		App: a,
	}
}

func MigrateNow(driver *database.Driver, dir string) error {
	m, err := Migrate(*driver, dir)
	if err != nil {
		return err
	}
	return m.Up()
}

func Migrate(databaseInstance database.Driver, dir string) (*migrate.Migrate, error) {
	dir = "/Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations"
	return migrate.NewWithDatabaseInstance("file://"+dir, "ql", databaseInstance)
}
