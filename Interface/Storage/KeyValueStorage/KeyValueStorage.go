package KeyValueStorage

import (
	"time"
)

type KeyValueStorage interface {
	Set(key string, value interface{}, expiration time.Duration) error
	MustSet(key string, value interface{}, expiration time.Duration)

	Get(key string) (ValueWrap, error)
	MustGet(key string) ValueWrap
}
