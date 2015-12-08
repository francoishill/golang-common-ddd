package HttpRequestHelper

import (
	"strconv"
	"strings"
)

type requestValueType int

const (
	RequestValue_Param requestValueType = iota
	RequestValue_Query
)

type RequestValue interface {
	HasValue() bool
	Type() requestValueType

	Bool() bool

	Float32() float32
	Float64() float64

	Int() int
	Int16() int16
	Int32() int32
	Int64() int64
	Int8() int8

	String() string

	Uint() uint
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64
	Uint8() uint8
}

func mustParseBool(s string) bool {
	intVal, err := strconv.ParseBool(s)
	if err != nil {
		trimmedStrLowerCase := strings.ToLower(strings.TrimSpace(s))
		switch {
		case trimmedStrLowerCase == "true":
			return true
		case trimmedStrLowerCase == "false":
			return false
		case trimmedStrLowerCase == "1":
			return true
		case trimmedStrLowerCase == "0":
			return false
		case trimmedStrLowerCase == "yes":
			return true
		case trimmedStrLowerCase == "no":
			return false
		}
		panic(err)
	}
	return intVal
}

func mustParseInt64(s string) int64 {
	intVal, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return intVal
}

func mustParseUint64(s string) uint64 {
	uintVal, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return uintVal
}

func mustParseFloat64(s string) float64 {
	floatVal, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return floatVal
}

type requestValue struct {
	hasValue    bool
	typeOfValue requestValueType
	stringValue string
}

func (r *requestValue) HasValue() bool {
	return r.hasValue
}
func (r *requestValue) Type() requestValueType {
	return r.typeOfValue
}

func (r *requestValue) Bool() bool {
	if !r.hasValue {
		return false
	}
	return mustParseBool(r.stringValue)
}

func (r *requestValue) Float32() float32 {
	if !r.hasValue {
		return 0.0
	}
	return float32(mustParseFloat64(r.stringValue))
}
func (r *requestValue) Float64() float64 {
	if !r.hasValue {
		return 0.0
	}
	return mustParseFloat64(r.stringValue)
}

func (r *requestValue) Int() int {
	if !r.hasValue {
		return 0
	}
	return int(mustParseInt64(r.stringValue))
}
func (r *requestValue) Int16() int16 {
	if !r.hasValue {
		return 0
	}
	return int16(mustParseInt64(r.stringValue))
}
func (r *requestValue) Int32() int32 {
	if !r.hasValue {
		return 0
	}
	return int32(mustParseInt64(r.stringValue))
}
func (r *requestValue) Int64() int64 {
	if !r.hasValue {
		return 0
	}
	return mustParseInt64(r.stringValue)
}
func (r *requestValue) Int8() int8 {
	if !r.hasValue {
		return 0
	}
	return int8(mustParseInt64(r.stringValue))
}

func (r *requestValue) String() string {
	if !r.hasValue {
		return ""
	}
	return r.stringValue
}

func (r *requestValue) Uint() uint {
	if !r.hasValue {
		return 0
	}
	return uint(mustParseUint64(r.stringValue))
}
func (r *requestValue) Uint16() uint16 {
	if !r.hasValue {
		return 0
	}
	return uint16(mustParseUint64(r.stringValue))
}
func (r *requestValue) Uint32() uint32 {
	if !r.hasValue {
		return 0
	}
	return uint32(mustParseUint64(r.stringValue))
}
func (r *requestValue) Uint64() uint64 {
	if !r.hasValue {
		return 0
	}
	return mustParseUint64(r.stringValue)
}
func (r *requestValue) Uint8() uint8 {
	if !r.hasValue {
		return 0
	}
	return uint8(mustParseUint64(r.stringValue))
}

func NewRequestValue(hasValue bool, typeOfValue requestValueType, stringValue string) RequestValue {
	return &requestValue{
		hasValue,
		typeOfValue,
		stringValue,
	}
}
