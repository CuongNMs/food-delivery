package uploadprovider

import (
	"context"
	"food-delivery/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session
	return provider
}

func (s3Provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	return nil, nil
}
