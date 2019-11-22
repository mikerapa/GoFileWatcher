package cli

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"github.com/radovskyb/watcher"
)


func DisplayWatchedFolderList(w *watcher.Watcher){
	// Print a list of all of the folders currently being watched.
	fmt.Println(Colorize("Watching the following folders:", GreenFg))
	for path, f := range w.WatchedFiles() {
		if f.IsDir() {
			fmt.Println(Colorize(path, WhiteFg))
		}
	}
}

func DisplayEvent(event watcher.Event) {
	switch event.Op {
	case watcher.Create:
		fmt.Println(Colorize(event.String(), GreenFg))
	case watcher.Remove:
		fmt.Println(Colorize(event.String(), RedFg))
	}
}