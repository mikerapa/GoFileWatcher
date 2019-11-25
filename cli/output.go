package cli

import (
	"github.com/logrusorgru/aurora"
	. "github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/radovskyb/watcher"
	"log"
	"os"
)
// colorizer
var au aurora.Aurora

func init() {
	// create colorizer and set the output of log
	au = aurora.NewAurora(isatty.IsTerminal(os.Stdout.Fd()))
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(colorable.NewColorableStdout())
}

func DisplayWatchedFolderList(w *watcher.Watcher){
	// Print a list of all of the folders currently being watched.
	log.Println(Colorize("Watching the following folders:", GreenFg))
	for path, f := range w.WatchedFiles() {
		if f.IsDir() {
			log.Printf("%v", Colorize(path, WhiteFg))
		}
	}
}

func DisplayEvent(event watcher.Event) {
	switch event.Op {
	case watcher.Create:
		log.Println(Colorize(event.String(), GreenFg))
	case watcher.Remove:
		log.Println(Colorize(event.String(), RedFg))
	case watcher.Rename, watcher.Chmod, watcher.Move, watcher.Write:
		log.Println(Colorize(event.String(), WhiteFg))
	}
}