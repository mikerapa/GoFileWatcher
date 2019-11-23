package main

import (
	"GoFileWatcher/cli"
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"os"
	"time"
)

func main() {
	commandLineSettings, err := cli.GetCommandLineSettings(os.Args[1:])
	if err!=nil {
		log.Fatal(err)
		return
	}


	w := watcher.New()

	go func() {
		for {
			select {
			case event := <-w.Event:
				// Print out the event to the screen.
				cli.DisplayEvent(event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// add watchers
	for _, folderPath:= range commandLineSettings.FolderPaths{
		if commandLineSettings.Recursive {
			if watcherAddError:= w.AddRecursive(folderPath); watcherAddError!= nil {
				log.Fatal(watcherAddError)
			}
		} else {
			if watcherAddError:= w.Add(folderPath); watcherAddError!= nil {
				log.Fatal(watcherAddError)
			}
		}

	}

	cli.DisplayWatchedFolderList(w)
	fmt.Println()


	// Start the watcher
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
