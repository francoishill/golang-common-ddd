package AutoReloadingSettings

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Settings/ReloadableSettings"
)

func (s *settings) AfterReloaded() {
	for _, h := range s.reloadObservers {
		h.OnSettingsReloaded()
	}
}

func (s *settings) SubscribeReloadObserver(handler SettingsReloadObserver) {
	s.reloadEventNotifyLock.Lock()
	defer s.reloadEventNotifyLock.Unlock()
	s.reloadObservers = append(s.reloadObservers, handler)
}

func (s *settings) UnsubscribeReloadObserver(handler SettingsReloadObserver) {
	s.reloadEventNotifyLock.Lock()
	defer s.reloadEventNotifyLock.Unlock()

	newObserverList := []SettingsReloadObserver{}
	for ind, _ := range s.reloadObservers {
		if s.reloadObservers[ind] != handler {
			newObserverList = append(newObserverList, handler)
		}
	}
	s.reloadObservers = newObserverList
}
