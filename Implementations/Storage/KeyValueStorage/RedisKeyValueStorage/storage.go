package RedisKeyValueStorage

import (
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"gopkg.in/redis.v3"
	"time"

	. "github.com/francoishill/golang-common-ddd/Interface/Storage/KeyValueStorage"
)

type storage struct {
	rc *redis.Client
}

func (s *storage) Set(key string, value interface{}, expiration time.Duration) error {
	return s.rc.Set(key, value, expiration).Err()
}
func (s *storage) MustSet(key string, value interface{}, expiration time.Duration) {
	err := s.Set(key, value, expiration)
	CheckError(err)
}

func (s *storage) Get(key string) (ValueWrap, error) {
	val := s.rc.Get(key)
	err := val.Err()
	if err != nil {
		return nil, err
	}
	return newValWrap(val), nil
}

func (s *storage) MustGet(key string) ValueWrap {
	val, err := s.Get(key)
	CheckError(err)
	return val
}

func New(redisOptions *redis.Options) KeyValueStorage {
	redisClient := redis.NewClient(redisOptions)
	err := redisClient.Ping().Err()
	CheckError(err)
	return &storage{redisClient}
}
