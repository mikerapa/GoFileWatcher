package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
)

type CommandLineSettings struct {
	FolderPaths []string
	Recursive   bool
}

func GetCommandLineSettings(commandLineArgs []string) (settings CommandLineSettings, err error){
	app := kingpin.New("GoFileWatcher", "File Watcher")
	var (
		pathsString = app.Flag("paths", "List of paths separated by semicolons.").Short('p').Required().String()
		recursive = app.Flag("recursive", "By default, the watcher is recursive. Use --no-recursive to make the watcher non-recursive.").Short('r').Default("true").Bool()
	)

	_, err = app.Parse(trimArray(commandLineArgs))

	// if parsing was successful, set the CommandLineSettings values
	if err == nil{
		settings = CommandLineSettings{FolderPaths: trimArray(parsePathArray(*pathsString)), Recursive:*recursive}
	}

	return
}

func parsePathArray(pathsString string) (paths []string) {
	paths = strings.Split(pathsString, ";")
	return
}

func trimArray(inputArray []string) (outputArray []string){
	for _ ,stringValue := range inputArray{
		trimmedValue := strings.Trim(stringValue, " ")
		if len(trimmedValue) >0{
			outputArray=append(outputArray, trimmedValue)
		}
	}
	return
}

