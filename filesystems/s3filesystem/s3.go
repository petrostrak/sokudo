package s3filesystem

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
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

	return listing, nil
}

func (s *S3) Delete(itemsToDelete []string) bool {
	return false
}

func (s *S3) Get(destination string, items ...string) error {
	return nil
}
