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
		log.Fatal(err) // TODO: Replace with another way to show error message
		return
	}


	w := watcher.New()

	go func() {
		for {
			select {
			case event := <-w.Event:
				//fmt.Println(event) // Print the event's info.
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
		if watcherAddError:= w.Add(folderPath); watcherAddError!= nil {
			log.Fatal(watcherAddError)
		}
	}

	cli.DisplayWatchedFolderList(w)
	fmt.Println()


	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
