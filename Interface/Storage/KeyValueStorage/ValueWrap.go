package KeyValueStorage

type ValueWrap interface {
	String() (string, error)
	MustString() string

	Bytes() ([]byte, error)
	MustBytes() []byte

	Float64() (float64, error)
	MustFloat64() float64

	Int64() (int64, error)
	MustInt64() int64

	Uint64() (uint64, error)
	MustUint64() uint64

	Scan(val interface{}) error
	MustScan(val interface{})
}
