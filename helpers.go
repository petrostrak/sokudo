package sokudo

import (
	"crypto/rand"
	"os"
)

const (
	randomString = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0987654321_+"
)

// RandomString generates a random string length n from values in
// const randomString.
func (s *Sokudo) RandomString(n int) string {
	str, r := make([]rune, n), []rune(randomString)

	for i := range str {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		str[i] = r[x%y]
	}

	return string(str)
}

func (s *Sokudo) CreateDirIfNotExists(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sokudo) CreateFileIfNotExists(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}
