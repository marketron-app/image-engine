package uploaders

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var (
	s3Config  aws.Config
	s3Session *session.Session
	s3Client  *s3.S3
)

func NewS3() {
	s3Config = aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewEnvCredentials(),
	}
	s3Session, _ = session.NewSession(&s3Config)
	s3Client = s3.New(s3Session)
}

func UploadToS3(path string, body []byte) error {

	// Upload the image to S3
	uploadParams := s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(path),
		Body:   bytes.NewReader(body),
	}
	_, err := s3Client.PutObject(&uploadParams)
	if err != nil {
		return err
	}
	return nil
}
