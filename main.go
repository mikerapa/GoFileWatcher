package main

import (
	"GoFileWatcher/cli"
	"fmt"
	FolderWatcher "github.com/mikerapa/FolderWatcher"
	"log"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	commandLineSettings, err := cli.GetCommandLineSettings(os.Args[1:])
	if err != nil {
		log.Fatal(err)
		return
	}

	watcher := FolderWatcher.New()
	paused := false

	pauseChannel := make(chan bool)
	exitChannel := make(chan bool)

	// add watchers from the command line
	for _, folderPath := range commandLineSettings.FolderPaths {
		// TODO the showHidden parameter is hard-coded
		if err := watcher.AddFolder(folderPath, commandLineSettings.Recursive, false); err != nil {
			// Just display the error and move on
			cli.DisplayError(err)
		}
	}

	// handle events
	go func() {
		for {
			select {

			case event := <-watcher.FileChanged:
				// Print out the event to the screen.
				if !paused {
					cli.DisplayEvent(event)
				}

			case <-watcher.Stopped:
				return

			case p := <-pauseChannel:
				// React to changes in the paused status
				if paused != p {
					paused = p
					cli.DisplayEventPause(paused)
				}
			case <-exitChannel:
				wg.Done()

			}
		}
	}()

	go cli.RunMenu(pauseChannel, exitChannel, &watcher)

	cli.DisplayWatchedFolderList(watcher.RequestedWatches)
	fmt.Println()

	// Start the watcher
	wg.Add(1)
	watcher.Start()

	// Shut down
	wg.Wait()
	println("Shutting down the watcher")
	watcher.Stop()
}

