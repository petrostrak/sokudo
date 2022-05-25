package webdavfilesystem

import "github.com/petrostrak/sokudo/filesystems"

type WebDAV struct {
	Host string
	User string
	Pass string
}

func (w *WebDAV) Put(fileName, folder string) error {
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
