package sokudo

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gabriel-vasile/mimetype"
	"github.com/petrostrak/sokudo/filesystems"
)

// For the last parameter fs, we want to pass a pointer to a FS if we want to upload to any
// remote FS(minio, s3, webDav or sftp) or nil for uploading to the local FS on the server.
func (s *Sokudo) UploadFile(r *http.Request, dst, field string, fs filesystems.FS) error {
	fileName, err := s.getFileToUpload(r, field)
	if err != nil {
		s.ErrorLog.Println(err)
		return err
	}

	if fs != nil {
		if err = fs.Put(fileName, dst); err != nil {
			s.ErrorLog.Println(err)
			return err
		}
	} else {
		// upload to the local filesystem on the server
		if err = os.Rename(fileName, fmt.Sprintf("%s/%s", dst, path.Base(fileName))); err != nil {
			s.ErrorLog.Println(err)
			return err
		}
	}

	return nil
}

func (s *Sokudo) getFileToUpload(r *http.Request, fieldName string) (string, error) {
	err := r.ParseMultipartForm(10 << 22)
	if err != nil {
		return "", err
	}

	file, header, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimeType, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}

	// go back to start of file
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	validMimeTypes := []string{
		"image/gif",
		"image/jpeg",
		"image/png",
		"application/pdf",
	}

	if !inSlice(validMimeTypes, mimeType.String()) {
		return "", errors.New("invalid file type uploaded")
	}

	dst, err := os.Create(fmt.Sprintf("./tmp/%s", header.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("./tmp/%s", header.Filename), nil
}

func inSlice(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
