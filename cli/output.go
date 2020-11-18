package cli

import (
	. "github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/mikerapa/FolderWatcher"
	"log"
)

func init() {
	// create colorize and set the output of log
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(colorable.NewColorableStdout())
}

func recursiveBoolToString(recursive bool) string {
	if recursive {
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

func DisplayUserMessage(message string) {
	log.Println(Colorize(message, WhiteFg))
}

func DisplayWatchedFolderList(folderList map[string]folderWatcher.WatchRequest) {
	// Print a list of all of the folders currently being watched.
	log.Println(Colorize("Watching the following folders:", GreenFg))
	for _, f := range folderList {
		log.Printf("%v (%v)", Colorize(f.Path, WhiteFg), Colorize(recursiveBoolToString(f.Recursive), YellowFg))
	}
	// TODO this does not show the showHidden value
}

func DisplayEventPause(pauseStatus bool) {
	if pauseStatus {
		log.Println(Colorize("Events", WhiteFg), Colorize("Paused", RedFg))
	} else {
		log.Println(Colorize("Events", WhiteFg), Colorize("Resumed", GreenFg))

	}
}

func DisplayEvent(event folderWatcher.FileEvent) {
	displayText := Sprintf("%s - %s", event.FileChange, event.Description)
	switch event.FileChange {

	case folderWatcher.Add:
		log.Println(Colorize(displayText, GreenFg))
	case folderWatcher.Remove:
		log.Println(Colorize(displayText, RedFg))
	case folderWatcher.Write:
		log.Println(Colorize(displayText, WhiteFg))
	case folderWatcher.Move:
		log.Println(Colorize(displayText, BlueFg))
	}
}
