package mininfilesystem

import (
	"context"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/petrostrak/sokudo/filesystems"
)

type Minio struct {
	Endpoint string
	Key      string
	Secret   string
	UseSSL   bool
	Region   string
	Bucket   string
}

func (m *Minio) getCredentials() *minio.Client {
	client, err := minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.Key, m.Secret, ""),
		Secure: m.UseSSL,
	})
	if err != nil {
		log.Println(err)
	}

	return client
}

func (m *Minio) Put(fileName, folder string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	objName := path.Base(fileName)
	client := m.getCredentials()

	uploadInfo, err := client.FPutObject(ctx, m.Bucket, fmt.Sprintf("%s/%s", folder, objName), fileName, minio.PutObjectOptions{})
	if err != nil {
		log.Println("Failed with FPutObject")
		log.Println(err)
		log.Println("Upload info:", uploadInfo)
		return err
	}

	return nil
}

func (m *Minio) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := m.getCredentials()

	objCh := client.ListObjects(ctx, m.Bucket, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	for obj := range objCh {
		if obj.Err != nil {
			fmt.Println(obj.Err)

			return listing, obj.Err
		}

		if !strings.HasPrefix(obj.Key, ".") {
			b := float64(obj.Size)
			kb := b / 1024
			mb := kb / 1024

			item := filesystems.Listing{
				Etag:         obj.ETag,
				LastModified: obj.LastModified,
				Key:          obj.Key,
				Size:         mb,
			}

			listing = append(listing, item)
		}
	}

	return listing, nil
}
