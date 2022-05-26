package sftpfilesystem

import (
	"fmt"
	"log"

	"github.com/petrostrak/sokudo/filesystems"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFTP struct {
	Host string
	User string
	Pass string
	Port string
}

func (s *SFTP) getCredentials() (*sftp.Client, error) {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)

	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	cwd, err := client.Getwd()
	if err != nil {
		return nil, err
	}

	log.Println("current working dir", cwd)

	return client, nil
}

func (s *SFTP) Put(fileName, folder string) error {
	return nil
}

func (s *SFTP) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing

	return listing, nil
}

func (s *SFTP) Delete(itemsToDelete []string) bool {
	return false
}

func (s *SFTP) Get(destination string, items ...string) error {
	return nil
}
