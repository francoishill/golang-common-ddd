package Errors

type ErrorsService interface {
	PanicClientErrorLocal(e interface{})
	PanicClientErrorLocal_FormattedString(format string, args ...interface{})
}
