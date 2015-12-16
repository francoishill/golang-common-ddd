package RedisKeyValueStorage

import (
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"gopkg.in/redis.v3"

	. "github.com/francoishill/golang-common-ddd/Interface/Storage/KeyValueStorage"
)

type valwrap struct {
	c *redis.StringCmd
}

func (v *valwrap) String() (string, error) { return v.c.Val(), nil }
func (v *valwrap) MustString() string {
	s := v.c.Val()
	return s
}

func (v *valwrap) Bytes() ([]byte, error) { return v.c.Bytes() }
func (v *valwrap) MustBytes() []byte {
	b, err := v.c.Bytes()
	CheckError(err)
	return b
}

func (v *valwrap) Float64() (float64, error) { return v.c.Float64() }
func (v *valwrap) MustFloat64() float64 {
	f, err := v.c.Float64()
	CheckError(err)
	return f
}

func (v *valwrap) Int64() (int64, error) { return v.c.Int64() }
func (v *valwrap) MustInt64() int64 {
	i, err := v.c.Int64()
	CheckError(err)
	return i
}

func (v *valwrap) Uint64() (uint64, error) { return v.c.Uint64() }
func (v *valwrap) MustUint64() uint64 {
	u, err := v.c.Uint64()
	CheckError(err)
	return u
}

func (v *valwrap) Scan(val interface{}) error {
	return v.c.Scan(val)
}
func (v *valwrap) MustScan(val interface{}) {
	err := v.c.Scan(val)
	CheckError(err)
}

func newValWrap(stringCmd *redis.StringCmd) ValueWrap {
	return &valwrap{
		stringCmd,
	}
}
