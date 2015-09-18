package Logger

//RFC 5424
type Logger interface {
	Emergency(msg string, params ...interface{})
	Alert(msg string, params ...interface{})
	Critical(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Notice(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Debug(msg string, params ...interface{})
}
