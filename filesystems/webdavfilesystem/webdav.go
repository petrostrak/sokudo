package webdavfilesystem

import (
	"fmt"
	"os"
	"path"

	"github.com/petrostrak/sokudo/filesystems"
	"github.com/studio-b12/gowebdav"
)

type WebDAV struct {
	Host string
	User string
	Pass string
}

func (w *WebDAV) getCredentials() *gowebdav.Client {
	return gowebdav.NewClient(w.Host, w.User, w.Pass)
}

func (w *WebDAV) Put(fileName, folder string) error {
	client := w.getCredentials()

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = client.WriteStream(fmt.Sprintf("%s/%s", folder, path.Base(fileName)), file, 0664)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebDAV) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing

	return listing, nil
}

func (w *WebDAV) Delete(itemsToDelete []string) bool {
	return false
}

func (w *WebDAV) Get(destination string, items ...string) error {
	return nil
}
