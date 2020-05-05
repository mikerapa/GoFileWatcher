package main

import (
	"GoFileWatcher/cli"
	"GoFileWatcher/fileWatcher"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	commandLineSettings, err := cli.GetCommandLineSettings(os.Args[1:])
	if err != nil {
		log.Fatal(err)
		return
	}

	watchMan := fileWatcher.NewWatchManager()
	paused := false

	pauseChannel := make(chan bool)
	exitChannel := make(chan bool)
	addChannel := make(chan fileWatcher.Folder)

	go func() {
		for {
			select {
			case folderToAdd := <-addChannel:
				watchMan.AddFolder(folderToAdd.Path, folderToAdd.Recursive)
				cli.DisplayFolderAdded(folderToAdd.Path, folderToAdd.Recursive)
			case event := <-watchMan.Watcher.Event:
				// Print out the event to the screen.
				if !paused {
					cli.DisplayEvent(event)
				}
			case err := <-watchMan.Watcher.Error:
				log.Fatalln(err)
			case <-watchMan.Watcher.Closed:
				return

			case p := <-pauseChannel:
				paused = p
			case <-exitChannel:
				ExitApplication(watchMan)
			}
		}
	}()

	go cli.RunMenu2(pauseChannel, exitChannel, addChannel, len(commandLineSettings.FolderPaths) == 0)

	// add watchers from the command line
	for _, folderPath := range commandLineSettings.FolderPaths {
		watchMan.AddFolder(folderPath, commandLineSettings.Recursive)
	}

	cli.DisplayWatchedFolderList(watchMan.FolderList)
	fmt.Println()

	// Start the watcher
	if err := watchMan.Watcher.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

//Shut down the application
func ExitApplication(watchMan *fileWatcher.WatchManager) {
	println("Shutting down File Watcher")
	watchMan.Close()
	os.Exit(0)
}
