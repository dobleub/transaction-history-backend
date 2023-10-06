package aws

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/private/protocol/rest"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dobleub/transaction-history-backend/internal/config"
)

func initAWSConnection(config *config.AWSConfig) (*s3.S3, error) {
	creds := credentials.NewStaticCredentials(config.AccessKeyId, config.SecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(config.DefaultRegion).WithCredentials(creds)

	svc := s3.New(session.New(), cfg)
	return svc, nil
}

func S3UploadDecoded(config *config.AWSConfig, b64, path string) (*s3.PutObjectOutput, error) {
	decode, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	svc, err := initAWSConnection(config)
	if err != nil {
		return nil, err
	}
	fileBytes := bytes.NewReader(decode)
	fileType := http.DetectContentType(decode)
	params := &s3.PutObjectInput{
		Bucket:        aws.String(config.Bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(int64(len(decode))),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func S3UploadNotDecoded(config *config.AWSConfig, fileDir string, buffer *bytes.Buffer) error {
	svc, err := initAWSConnection(config)
	if err != nil {
		return err
	}
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(config.Bucket),
		Key:                  aws.String(fileDir),
		Body:                 bytes.NewReader(buffer.Bytes()),
		ContentLength:        aws.Int64(int64(buffer.Len())),
		ContentType:          aws.String("text/csv"),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		ACL:                  aws.String("public-read"),
	})
	if err != nil {
		return err
	}
	return nil
}

func GetURLObject(config *config.AWSConfig, fileDir string) (string, error) {
	svc, err := initAWSConnection(config)
	if err != nil {
		return "", err
	}
	getObj, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(fileDir),
	})
	if getObj == nil {
		return "", fmt.Errorf("AWS: S3 object not found")
	}
	rest.Build(getObj)
	url := getObj.HTTPRequest.URL.String()
	return url, nil
}

func DeleteObject(config *config.AWSConfig, fileDir string) error {
	svc, err := initAWSConnection(config)
	if err != nil {
		return err
	}
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(fileDir),
	})
	if err != nil {
		return err
	}
	return nil
}

func DownloadObject(config *config.AWSConfig, fileDir string) (*bytes.Buffer, error) {
	svc, err := initAWSConnection(config)
	if err != nil {
		return nil, err
	}
	getObj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(fileDir),
	})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(getObj.Body)
	return buf, nil
}
