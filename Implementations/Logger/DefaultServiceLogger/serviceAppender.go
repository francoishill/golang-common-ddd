package DefaultServiceLogger

import (
	"github.com/ayufan/golang-kardianos-service"
	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/levels"
)

type serviceAppender struct {
	currentLayout layout.Layout
	svcLogger     service.Logger
}

func (s *serviceAppender) Layout() layout.Layout {
	return s.currentLayout
}

func (s *serviceAppender) SetLayout(l layout.Layout) {
	s.currentLayout = l
}

func (s *serviceAppender) Write(level levels.LogLevel, message string, args ...interface{}) {
	if level == levels.FATAL || level == levels.ERROR {
		s.svcLogger.Errorf(message, args...)
	} else if level == levels.WARN {
		s.svcLogger.Warningf(message, args...)
	} else {
		s.svcLogger.Infof(message, args...)
	}
}

func newAppender(layout layout.Layout, svcLogger service.Logger) appenders.Appender {
	return &serviceAppender{
		layout,
		svcLogger,
	}
}
