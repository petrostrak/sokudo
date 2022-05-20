package cache

import (
	"time"

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
	var fromCache []byte

	err := b.Conn.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(s))
		if err != nil {
			return err
		}

		if err = item.Value(func(val []byte) error {
			fromCache = append(fromCache, val...)
			return nil
		}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	decoded, err := decode(string(fromCache))
	if err != nil {
		return nil, err
	}

	item := decoded[s]

	return item, nil
}

func (b *BadgerCache) Set(s string, value interface{}, expires ...int) error {
	entry := Entry{}
	entry[s] = value
	encoded, err := encode(entry)
	if err != nil {
		return err
	}

	if len(expires) > 0 {
		err = b.Conn.Update(func(txn *badger.Txn) error {
			e := badger.NewEntry([]byte(s), encoded).WithTTL(time.Second * time.Duration(expires[0]))
			err = txn.SetEntry(e)
			return err
		})
	} else {
		err = b.Conn.Update(func(txn *badger.Txn) error {
			e := badger.NewEntry([]byte(s), encoded)
			err = txn.SetEntry(e)
			return err
		})
	}

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
