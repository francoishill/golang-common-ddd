package DefaultLogger

import (
	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/levels"
)

//
//TODO: Remove this file when PULL REQUEST ACCEPTED -- https://github.com/ian-kent/go-log
//
type multipleAppender struct {
	currentLayout   layout.Layout
	listOfAppenders []appenders.Appender
}

func (m *multipleAppender) Layout() layout.Layout {
	return m.currentLayout
}

func (m *multipleAppender) SetLayout(l layout.Layout) {
	m.currentLayout = l
}

func (m *multipleAppender) Write(level levels.LogLevel, message string, args ...interface{}) {
	for _, appender := range m.listOfAppenders {
		appender.Write(level, message, args...)
	}
}

func Multiple(layout layout.Layout, appenders ...appenders.Appender) appenders.Appender {
	return &multipleAppender{
		listOfAppenders: appenders,
		currentLayout:   layout,
	}
}
