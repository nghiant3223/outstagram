package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisSupplier struct {
	*redis.Client
}

var supplier *RedisSupplier

func NewRedisSupplier() (*RedisSupplier, error) {
	if supplier == nil {
		supplier = &RedisSupplier{}
		supplier.Client = redis.NewClient(&redis.Options{
			Addr:     "cache:6379",
			Password: "",
			DB:       0,
		})
	}

	if _, err := supplier.Client.Ping().Result(); err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *RedisSupplier) Get(key string) *redis.StringCmd {
	return s.Client.Get(key)
}

func (s *RedisSupplier) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return s.Client.Set(key, value, expiration)
}

func (s *RedisSupplier) save(key string, value interface{}, expiry time.Duration) error {
	if bytes, err := GetBytes(value); err != nil {
		return err
	} else {
		if err := s.Client.Set(key, bytes, expiry).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (s *RedisSupplier) load(key string, writeTo interface{}) (bool, error) {
	if data, err := s.Client.Get(key).Bytes(); err != nil {
		if err == redis.Nil {
			return false, nil
		} else {
			return false, err
		}
	} else {
		if err := DecodeBytes(data, writeTo); err != nil {
			return false, err
		}
	}
	return true, nil
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeBytes(input []byte, thing interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(input))
	return dec.Decode(thing)
}

func (s *RedisSupplier) MigratePostFeedToCache(userID uint, postIDs []uint) error {
	return s.save(fmt.Sprint(userID), postIDs, 0)
}

func (s *RedisSupplier) GetPostFeed(userID uint, postIDs []uint) (bool, error) {
	return s.load(fmt.Sprint(userID), postIDs)
}

func (s *RedisSupplier) UpdatePostFeed(userID uint, postID uint) error {
	var postIDs []uint

	_, err := s.GetPostFeed(userID, postIDs)
	if err != nil {
		return err
	}

	postIDs = append([]uint{postID}, postIDs...)
	if err := s.save(fmt.Sprint(userID), postIDs, 0); err != nil {
		return err
	}

	return nil
}
