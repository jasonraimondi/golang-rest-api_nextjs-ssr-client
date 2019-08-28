package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"

	"git.jasonraimondi.com/jason/jasontest/app/lib/config"

	"git.jasonraimondi.com/jason/jasontest/app/lib/repository"
	"git.jasonraimondi.com/jason/jasontest/app/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type BucketName string

type PhotoUploadService struct {
	originals      BucketName
	repository     *repository.Factory
	userRepository *repository.UserRepository
	s3             *config.S3Config
}

func (s *PhotoUploadService) FileUpload(form *multipart.Form, userId string) *echo.HTTPError {
	files := form.File["files[]"]

	user, err := s.userRepository.GetById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	newSession, _ := session.NewSession(s.s3.Config)
	s3Client := s3.New(newSession)

	var errs []error
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			errs = append(errs, err)
		}
		size := fileHeader.Size
		buffer := make([]byte, size)
		_, _ = file.Read(buffer)

		photo, err := s.createPhoto(user, file, fileHeader)
		if err != nil {
			errs = append(errs, err)
		}
		put := &s3.PutObjectInput{
			ACL:                aws.String("public-read"),
			ContentDisposition: aws.String("attachment"),
			ContentLength:      aws.Int64(int64(photo.FileSize)),
			ContentType:        aws.String(http.DetectContentType(buffer)),
			Body:               bytes.NewReader(buffer),
			Bucket:             aws.String(string(s.originals)),
			Key:                aws.String(photo.RelativeURL),
		}
		if _, err = s3Client.PutObject(put); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return echo.NewHTTPError(http.StatusPartialContent, fmt.Sprintf("%d errors out of %d uploaded", len(errs), len(files)))
	}
	return nil
}

func (s *PhotoUploadService) createPhoto(user *models.User, f multipart.File, fileHeader *multipart.FileHeader) (photo *models.Photo, err error) {
	fileSHA256, _ := GetFileSHA256(f)
	photo = models.NewPhoto(
		uuid.NewV4(),
		user,
		fileHeader.Filename,
		fileSHA256,
		fileHeader.Header.Get("Content-Type"),
		uint64(fileHeader.Size),
	)
	if err := s.repository.PhotoRepository().Create(photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func GetFileSHA256(file io.Reader) (result string, err error) {
	hash := sha256.New()
	if _, err = io.Copy(hash, file); err != nil {
		return result, err
	}
	hashInBytes := hash.Sum(nil)
	result = hex.EncodeToString(hashInBytes)
	return result, nil
}
