package storage
import (
	"gopkg.in/fsnotify.v1"
	"log"
	"aurhobas/config"
	"github.com/joernweissenborn/stream2go"
	"strings"
)

func FileWatcher(events stream2go.StreamController){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.Contains(event,META_FILE_NAME) {
					if event.Op&fsnotify.Write == fsnotify.Write {
						events.Add(UpdateItem{event.Name})
					} else if event.Op&fsnotify.Create == fsnotify.Create{
						events.Add(NewItem{event.Name})
					} else if event.Op&fsnotify.Remove == fsnotify.Remove {
						events.Add(RemoveItem{event.Name})
					}
				}
			case err := <-watcher.Errors:
				log.Fatal("error:", err)
			}
		}
	}()

	err = watcher.Add(config.CFG.Root)
	if err != nil {
		log.Fatal(err)
	}
}