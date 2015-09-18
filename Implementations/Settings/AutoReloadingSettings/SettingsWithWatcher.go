package AutoReloadingSettings

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Settings/ReloadableSettings"
	"sync"
	"time"
)

type settings struct {
	filePathToWatch string

	handlers

	reloadEventNotifyLock *sync.RWMutex
	reloadObservers       []SettingsReloadObserver

	lastWatcherModifiedTime time.Time
}

func New(filePathToWatch string, handlers handlers) ReloadableSettings {
	s := &settings{
		filePathToWatch:       filePathToWatch,
		handlers:              handlers,
		reloadEventNotifyLock: &sync.RWMutex{},
	}
	s.watchConfig()
	return s
}
