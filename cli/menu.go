package cli

import (
	"GoFileWatcher/fileWatcher"
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func AddFolderMenu() (folderPath string, recursive bool, err error) {

	// Set up prompt for asking the user for a folder path
	validate := func(inputPath string) error {
		if !fileWatcher.IsValidDirPath(inputPath) {
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

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	_, result, err := recursivePrompt.Run()
	recursive = result != "Not Recursive"

	fmt.Printf("You choose %q\n", result)
	return
}

func RunMenu2(pauseChannel chan bool, exitChannel chan bool, addChannel chan fileWatcher.Folder, nothingToWatch bool) {
	paused := false

	prompt := promptui.Select{
		Label: "Main menu",
		Items: []string{"Add folder", "List folders", "Remove folder", "Resume watch", "Exit"},
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		if nothingToWatch {
			println("No paths are being watched")
		} else {
			if _, err := reader.ReadString('\n'); err != nil {
				DisplayError(err)
			}
			nothingToWatch = false
			paused = true
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Add folder":
			folderPath, recursive, err := AddFolderMenu()
			if err != nil {
				DisplayError(err)
			}
			addChannel <- fileWatcher.Folder{Path: folderPath, Recursive: recursive}
		case "List folders":
		case "Remove folder":
		case "Resume watch":
			paused = false
			pauseChannel <- false
			DisplayEventPause(paused)
		case "Exit":
			exitChannel <- true
		default:

		}

	}
}
