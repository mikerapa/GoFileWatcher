package cli

import (
	"GoFileWatcher/fileWatcher"
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/mikerapa/FolderWatcher"
	"os"
)

func RemoveFolderMenu(folderList map[string]folderWatcher.WatchRequest) (folderPath string, err error) {
	var items []string

	for folder := range folderList {
		items = append(items, folder)
	}
	removeFolderPrompt := promptui.Select{
		Label: "Remove Folder",
		Items: items,
	}

	_, folderPath, err = removeFolderPrompt.Run()

	return
}

func AddFolderMenu() (folderPath string, recursive bool, err error) {

	// Set up prompt for asking the user for a folder path
	validate := func(inputPath string) error {
		if len(inputPath) > 0 && !fileWatcher.IsValidDirPath(inputPath) {
			return errors.New("invalid path")
		}
		return nil
	}

	pathPrompt := promptui.Prompt{Label: "Enter Path", Validate: validate}

	// set up recursivePrompt to ask user if the new folder will be recursive
	recursivePrompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Recursive", "Not Recursive"},
	}

	folderPath, err = pathPrompt.Run()
	// If the user entered an empty folder path, get out of the AddFolderMenu
	if len(folderPath) == 0 {
		return
	}
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	_, result, err := recursivePrompt.Run()
	recursive = result != "Not Recursive"

	fmt.Printf("You choose %q\n", result)
	return
}

func RunMenu(pauseChannel chan bool, exitChannel chan bool, watcher *folderWatcher.Watcher) {
	prompt := promptui.Select{
		Label: "Main menu",
		Items: []string{"Add folder", "List folders", "Remove folder", "Resume watch", "Exit"},
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		if len(watcher.RequestedWatches) == 0 {
			DisplayUserMessage("No paths are being watched")
		} else {
			if _, err := reader.ReadString('\n'); err != nil {
				DisplayError(err)
			}
			pauseChannel<-true
		}
		_, result, err := prompt.Run()
		if err != nil {
			DisplayError(err)
			return
		}

		switch result {
		case "Add folder":
			folderPath, recursive, err := AddFolderMenu()
			if err != nil {
				DisplayError(err) // if there's an error from AddFolderMenu, get out
				continue
			}
			// Make sure there is a folder path to add
			if len(folderPath) == 0 {
				DisplayUserMessage("No folders added")
				continue
			}
			// TODO the showHidden value is hard-coded
			err = watcher.AddFolder(folderPath, recursive, false)
			if err == nil {
				DisplayFolderAdded(folderPath, recursive)
			} else {
				DisplayError(err)
			}
		case "List folders":
			DisplayWatchedFolderList(watcher.RequestedWatches)
		case "Remove folder":
			folderPath, err := RemoveFolderMenu(watcher.RequestedWatches)
			if err == nil {
				err = watcher.RemoveFolder(folderPath, true)
				if err != nil {
					DisplayError(err)
				}
			} else {
				DisplayError(err)
			}
		case "Resume watch":
			pauseChannel <- false
		case "Exit":
			exitChannel <- true
		default:

		}

	}
}
