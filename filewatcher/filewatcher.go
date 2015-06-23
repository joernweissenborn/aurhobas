package filewatcher
import (
	"gopkg.in/fsnotify.v1"
	"log"
	"aurhobas/config"
	"github.com/joernweissenborn/stream2go"
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
				events.Add(event)
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