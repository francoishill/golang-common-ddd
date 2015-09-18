package ReloadableSettings

type ReloadableSettings interface {
	SubscribeReloadObserver(handler SettingsReloadObserver)
	UnsubscribeReloadObserver(handler SettingsReloadObserver)
}
