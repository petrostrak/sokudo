package cache

import (
	"github.com/dgraph-io/badger/v3"
)

type BadgerCache struct {
	Conn   *badger.DB
	Prefix string
}

func (b *BadgerCache) Has(s string) (bool, error) {
	_, err := b.Get(s)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (b *BadgerCache) Get(s string) (interface{}, error) {

	return nil, nil
}

func (b *BadgerCache) Set(s string, value interface{}, expires ...int) error {

	return nil
}

func (b *BadgerCache) Forget(s string) error {

	return nil
}

func (b *BadgerCache) EmptyByMatch(s string) error {

	return nil
}

func (b *BadgerCache) Empty() error {

	return nil
}
