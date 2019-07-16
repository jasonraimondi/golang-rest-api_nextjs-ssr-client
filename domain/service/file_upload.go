package service

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
//	tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
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
//		ContentType:          aws.String(http.DetectContentType(buffer)),
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
	tx := s.repository.DBx.MustBegin()

	newSession := session.New(s.s3Config)

	bucket := aws.String("wasabi-golang-example")
	key := aws.String("wasabi-testobject")

	s3Client := s3.New(newSession)

	for _, file := range files {
		photo := model.NewPhoto(
			uuid.NewV4(),
			user,
			file.Filename,
			file.Header.Get("Content-Type"),
			file.Size,
		)
		fofile, _ := file.Open()
		buffer := make([]byte, file.Size)
		fofile.Read(buffer)
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			ContentLength: aws.Int64(int64(file.Size)),
			ContentType:   aws.String(http.DetectContentType(buffer)),
			Body:          bytes.NewReader(buffer),
			Bucket:        bucket,
			Key:           key,
		})

		if err = repository.CreatePhotoTx(tx, photo); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
