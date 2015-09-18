package AutoReloadingSettings

import (
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"github.com/howeyc/fsnotify"
	"time"
)

func (s *settings) watchConfig() {
	watcher, err := fsnotify.NewWatcher()
	CheckError(err)

	go func() {
		for {
			select {
			case _ = <-watcher.Event: //case e := <-watcher.Event:
				if s.lastWatcherModifiedTime.Add(1 * time.Second).After(time.Now()) {
					continue
				}
				s.lastWatcherModifiedTime = time.Now()

				s.handlers.ValidateAndUseFile(s.filePathToWatch)
				s.AfterReloaded()
				break
			case err := <-watcher.Error:
				s.handlers.OnWatchReloadError(err)
				break
			}
		}
	}()

	err = watcher.Watch(s.filePathToWatch)
	CheckError(err)
}
