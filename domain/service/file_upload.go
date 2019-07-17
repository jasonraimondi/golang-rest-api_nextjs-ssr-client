package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	"git.jasonraimondi.com/jason/jasontest/domain/model"
	"git.jasonraimondi.com/jason/jasontest/domain/repository"
)

//func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
//	// get the file size and read
//	// the file content into a buffer
//	size := fileHeader.Size
//	buffer := make([]byte, size)
//	file.Read(buffer)
//
//	// create a unique file name for the file
//
//
//	// config settings: this is where you choose the bucket,
//	// filename, content-type and storage class of the file
//	// you're uploading
//	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
//		Bucket:               aws.String("test-bucket"),
//		Key:                  aws.String(tempFileName),
//		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
//		Body:                 bytes.NewReader(buffer),
//		ContentLength:        aws.Int64(int64(size)),
//		MimeType:          aws.String(http.DetectContentType(buffer)),
//		ContentDisposition:   aws.String("attachment"),
//		ServerSideEncryption: aws.String("AES256"),
//		StorageClass:         aws.String("INTELLIGENT_TIERING"),
//	})
//	if err != nil {
//		return "", err
//	}
//
//	return tempFileName, err
//}

func (s *Service) FileUpload(form *multipart.Form, userId string) *echo.HTTPError {
	files := form.File["file[]"]

	user, err := s.repository.User().GetById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	newSession, _ := session.NewSession(s.s3Config.Config)
	s3Client := s3.New(newSession)

	tx := s.repository.DBx.MustBegin()
	for _, fileHeader := range files {
		file, erroz := fileHeader.Open()
		if erroz != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, erroz)
		}
		size := fileHeader.Size
		buffer := make([]byte, size)
		_, _ = file.Read(buffer)

		photo, httpError := CreatePhoto(tx, user, file, fileHeader)
		if httpError != nil {
			return httpError
		}
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			ACL:                aws.String("public-read"),
			ContentDisposition: aws.String("attachment"),
			ContentLength:      aws.Int64(int64(photo.FileSize)),
			ContentType:        aws.String(http.DetectContentType(buffer)),
			Body:               bytes.NewReader(buffer),
			Bucket:             aws.String(s.s3Config.OriginalBucket),
			Key:                aws.String(photo.RelativeURL),
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}

func CreatePhoto(tx *sqlx.Tx, user *model.User, f multipart.File, fileHeader *multipart.FileHeader) (photo *model.Photo, httpError *echo.HTTPError) {
	fileSHA256, _ := GetFileSHA256(f)
	photo = model.NewPhoto(
		uuid.NewV4(),
		user,
		fileHeader.Filename,
		fileSHA256,
		fileHeader.Header.Get("Content-Type"),
		fileHeader.Size,
	)
	if err := repository.CreatePhotoTx(tx, photo); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
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
