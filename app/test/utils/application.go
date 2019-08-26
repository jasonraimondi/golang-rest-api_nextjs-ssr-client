package utils

import (
	"github.com/jinzhu/gorm"

	"git.jasonraimondi.com/jason/jasontest/app/lib"
	"git.jasonraimondi.com/jason/jasontest/app/lib/awsupload"
	"git.jasonraimondi.com/jason/jasontest/server/handlers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	_ "github.com/mattn/go-sqlite3"
)

func NewTestApplication(tables []interface{}) (a *lib.Application) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(tables...)
	sessionToken := ""
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("", "", sessionToken),
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
	}
	s3Config := awsupload.NewS3Config("test-originals", config)
	a = lib.NewApplication(
		db,
		s3Config,
		"jwtSecureKey-test",
		"/Users/jason/go/src/git.jasonraimondi.com/jason/jasontest/db/migrations",
	)
	return a
}

func NewTestHandler(tables []interface{}) *handlers.Handler {
	a := NewTestApplication(tables)
	return &handlers.Handler{
		App: a,
	}
}
