package cli

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func MainMenu() (exitApplication bool) {
	prompt := promptui.Select{
		Label: "Main menu",
		Items: []string{"Add folder", "List folders", "Remove folder", "Resume watch", "Exit"},
	}
	// TODO add menu option for pause and unpause
	//exitMenu := false
	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Add folder":
		case "List folders":
		case "Remove folder":
		case "Resume watch":
			return false
		case "Exit":
			return true
		default:

		}
		//fmt.Printf("You choose %q\n", result)
	}
	return
}

// Start the menu system and listen for the Enter key
func RunMenu(pauseChannel chan bool, exitChannel chan bool) {
	paused := false
	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadString('\n')
		paused = !paused
		pauseChannel <- paused
		DisplayEventPause(paused)

		if paused {
			exitApplication := MainMenu()
			if exitApplication {
				exitChannel <- true
			} else {
				// when exiting the main menu, anything but the exit application option should resume the file watching
				paused = false
				DisplayEventPause(paused)
			}
		}
	}
}
