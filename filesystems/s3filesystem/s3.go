package s3filesystem

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/petrostrak/sokudo/filesystems"
)

type S3 struct {
	Key      string
	Secret   string
	Region   string
	Endpoint string
	Bucket   string
}

func (s *S3) getCredentials() *credentials.Credentials {
	return credentials.NewStaticCredentials(s.Key, s.Secret, "")
}

func (s *S3) Put(fileName, folder string) error {
	return nil
}

func (s *S3) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing

	client := s.getCredentials()
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    &s.Endpoint,
		Region:      &s.Region,
		Credentials: client,
	}))

	svc := s3.New(sess)
	input := &s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(prefix),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(aerr.Error())
		}

		return nil, err
	}

	for _, key := range result.Contents {
		b := float64(*key.Size)
		kb := b / 1024
		mb := kb / 1024

		current := filesystems.Listing{
			Etag:         *key.ETag,
			LastModified: *key.LastModified,
			Key:          *key.Key,
			Size:         mb,
		}

		listing = append(listing, current)
	}

	return listing, nil
}

func (s *S3) Delete(itemsToDelete []string) bool {
	return false
}

func (s *S3) Get(destination string, items ...string) error {
	return nil
}
