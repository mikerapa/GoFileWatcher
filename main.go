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
	if err != nil {
		log.Fatal(err)
		return
	}

	w := watcher.New()

	paused := false

	pauseChannel := make(chan bool)
	exitChannel := make(chan bool)

	go func() {
		for {
			select {
			case event := <-w.Event:
				// Print out the event to the screen.
				if !paused {
					cli.DisplayEvent(event)
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			case p := <-pauseChannel:
				paused = p
				println("paused", paused)
			case <-exitChannel:
				ExitApplication(w)
			}
		}
	}()

	go cli.RunMenu(pauseChannel, exitChannel)
	//// keyboard listen loop
	//reader := bufio.NewReader(os.Stdin)
	//go func (reader *bufio.Reader, w *watcher.Watcher){
	//	for {
	//		text, _ := reader.ReadString('\n')
	//		if len(text)>0{
	//			println("text", text)
	//
	//		}
	//		paused = !paused
	//		cli.DisplayEventPause(paused)
	//		if paused{
	//			exitApplication := cli.MainMenu()
	//			if exitApplication{
	//				ExitApplication(w)
	//			} else {
	//				paused = false
	//				cli.DisplayEventPause(paused)
	//			}
	//		}
	//	}
	//}(reader, w)

	// add watchers
	for _, folderPath := range commandLineSettings.FolderPaths {
		if commandLineSettings.Recursive {
			if watcherAddError := w.AddRecursive(folderPath); watcherAddError != nil {
				log.Fatal(watcherAddError)
			}
		} else {
			if watcherAddError := w.Add(folderPath); watcherAddError != nil {
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

//Shut down the application
func ExitApplication(w *watcher.Watcher) {
	println("Shutting down File Watcher")
	w.Close()
	os.Exit(0)
}
