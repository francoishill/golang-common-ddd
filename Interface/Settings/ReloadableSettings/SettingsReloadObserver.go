package ReloadableSettings

type SettingsReloadObserver interface {
	OnSettingsReloaded()
}
