package uploadprovider

import (
	"bytes"
	"coffee_api/commons"
	"coffee_api/helpers"
	"coffee_api/modules/upload/entity"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName, region, apiKey, secret, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Sesson, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})

	if err != nil {
		helpers.Fatal(err)
	}

	provider.session = s3Sesson

	return provider
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, uploadDto *entity.UploadDTO) (*commons.Image, error) {
	fileBytes := bytes.NewReader(uploadDto.Data)
	fileType := http.DetectContentType(uploadDto.Data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(uploadDto.Dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	img := &commons.Image{
		Url:       uploadDto.Dst,
		CloudName: "s3",
	}

	fmt.Println("img:: ", img)

	return img, nil
}
