package DefaultLogger

import (
	"fmt"
	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/levels"
	golog "github.com/ian-kent/go-log/log"
	gologger "github.com/ian-kent/go-log/logger"
	"time"

	. "github.com/francoishill/golang-common-ddd/Interface/Logger"
)

type logger struct {
	logger gologger.Logger
}

func (l *logger) formatNowTime() string {
	now := time.Now()
	_, offset := now.Zone()
	var timezoneSign string
	if offset >= 0 {
		timezoneSign = "+"
	} else {
		timezoneSign = "-"
	}
	return fmt.Sprintf("%s %s%d", now.Local().Format("2006-01-02 15:04:05"), timezoneSign, offset/(60*60))
}

func (l *logger) combineParams(msg string, params ...interface{}) []interface{} {
	combined := []interface{}{
		fmt.Sprintf("{%s} %s", l.formatNowTime(), msg),
	}
	combined = append(combined, params...)
	return combined
}

func (l *logger) Emergency(msg string, params ...interface{}) {
	l.logger.Fatal(l.combineParams(msg, params...))
}
func (l *logger) Alert(msg string, params ...interface{}) {
	l.logger.Error(l.combineParams(msg, params...))
}
func (l *logger) Critical(msg string, params ...interface{}) {
	l.logger.Error(l.combineParams(msg, params...))
}
func (l *logger) Error(msg string, params ...interface{}) {
	l.logger.Error(l.combineParams(msg, params...))
}
func (l *logger) Warning(msg string, params ...interface{}) {
	l.logger.Warn(l.combineParams(msg, params...))
}
func (l *logger) Notice(msg string, params ...interface{}) {
	l.logger.Info(l.combineParams(msg, params...))
}
func (l *logger) Informational(msg string, params ...interface{}) {
	l.logger.Info(l.combineParams(msg, params...))
}
func (l *logger) Debug(msg string, params ...interface{}) {
	l.logger.Debug(l.combineParams(msg, params...))
}

func getPrefixWithSpace(prefix string) string {
	if prefix == "" {
		return ""
	} else {
		return prefix + " "
	}
}

func New(logFileName, prefix string, isDevMode bool) Logger {
	l := golog.Logger()

	if isDevMode {
		l.SetLevel(levels.TRACE)
	} else {
		l.SetLevel(levels.INFO)
	}

	layoutToUse := layout.Pattern(getPrefixWithSpace(prefix) + "[%p] %m") //level/priority, message

	rollingFileAppender := appenders.RollingFile(logFileName, true)
	rollingFileAppender.MaxBackupIndex = 5
	rollingFileAppender.MaxFileSize = 20 * 1024 * 1024 // 20 MB
	rollingFileAppender.SetLayout(layoutToUse)

	consoleAppender := appenders.Console()
	consoleAppender.SetLayout(layoutToUse)
	l.SetAppender(
		Multiple( //appenders.Multiple( ONCE PULL REQUEST OF ABOVE IS IN
			layoutToUse,
			rollingFileAppender,
			consoleAppender,
		))

	return &logger{l}
}
