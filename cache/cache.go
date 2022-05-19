package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int) error
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]interface{}

func encode(item Entry) ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	err := e.Encode(item)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func decode(s string) (Entry, error) {
	item := Entry{}
	b := bytes.Buffer{}
	b.Write([]byte(s))

	d := gob.NewDecoder(&b)

	err := d.Decode(&item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *RedisCache) Has(s string) (bool, error) {
	key := fmt.Sprintf("%s:%s", c.Prefix, s)
	conn := c.Conn.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (c *RedisCache) Get(s string) (interface{}, error) {

	return "", nil
}

func (c *RedisCache) Set(s string, data interface{}, ttl ...int) error {

	return nil
}

func (c *RedisCache) Forget(s string) error {

	return nil
}

func (c *RedisCache) EmptyByMatch(s string) error {

	return nil
}

func (c *RedisCache) Empty() error {

	return nil
}
