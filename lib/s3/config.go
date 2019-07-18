package s3

import (
	"github.com/aws/aws-sdk-go/aws"
)

type S3Config struct {
	OriginalBucket string
	*aws.Config
}

func NewS3Config(originalBucket string, aws *aws.Config) *S3Config {
	return &S3Config{
		OriginalBucket: originalBucket,
		Config:         aws,
	}
}

