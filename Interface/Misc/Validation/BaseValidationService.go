package Validation

type BaseValidationService interface {
	IsEmail(str string) bool
}
