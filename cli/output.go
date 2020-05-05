package cli

import (
	"GoFileWatcher/fileWatcher"
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

func recursiveBoolToString(recurisve bool) string {
	if recurisve {
		return "Recursive"
	} else {
		return "Not Recursive"
	}
}

func DisplayError(err error) {
	log.Println(Colorize("Error:", RedFg), Colorize(err.Error(), WhiteFg))
}

func DisplayFolderAdded(path string, recursive bool) {
	log.Printf("Now watching %v (%v)", Colorize(path, WhiteFg), Colorize(recursiveBoolToString(recursive), YellowFg))
}

func DisplayWatchedFolderList(folderList map[string]fileWatcher.Folder) {
	// Print a list of all of the folders currently being watched.
	log.Println(Colorize("Watching the following folders:", GreenFg))
	for _, f := range folderList {
		log.Printf("%v (%v)", Colorize(f.Path, WhiteFg), Colorize(recursiveBoolToString(f.Recursive), YellowFg))
	}
}

func DisplayEventPause(pauseStatus bool) {
	if pauseStatus {
		log.Println(Colorize("Events", WhiteFg), Colorize("Paused", RedFg))
	} else {
		log.Println(Colorize("Events", WhiteFg), Colorize("Resumed", GreenFg))

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
