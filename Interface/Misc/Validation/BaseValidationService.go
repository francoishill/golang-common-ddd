package Validation

type BaseValidationService interface {
	InRange(value, left, right float64) bool
	IsASCII(str string) bool
	IsAlpha(str string) bool
	IsAlphanumeric(str string) bool
	IsBase64(str string) bool
	IsCreditCard(str string) bool
	IsEmail(str string) bool
	IsEmptyStringOrWhitespace(str string) bool
	IsFilePath(str string) (bool, int)
	IsFloat(str string) bool
	IsHexadecimal(str string) bool
	IsHexcolor(str string) bool
	IsIP(str string) bool
	IsIPv4(str string) bool
	IsIPv6(str string) bool
	IsInt(str string) bool
	IsJSON(str string) bool
	IsLatitude(str string) bool
	IsLongitude(str string) bool
	IsNatural(value float64) bool
	IsNegative(value float64) bool
	IsNonNegative(value float64) bool
	IsNonPositive(value float64) bool
	IsNumeric(str string) bool
	IsPositive(value float64) bool
	IsWhole(value float64) bool
}
