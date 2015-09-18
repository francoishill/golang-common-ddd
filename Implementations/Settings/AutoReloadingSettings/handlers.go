package AutoReloadingSettings

type handlers interface {
	ValidateAndUseFile(filePath string)
	OnWatchReloadError(err error)
}
