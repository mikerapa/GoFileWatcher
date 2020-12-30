package cli

import (
	"flag"
	"strings"
)

type CommandLineSettings struct {
	FolderPaths []string
	Recursive   bool
}

func GetCommandLineSettings(commandLineArgs []string) (settings CommandLineSettings, err error) {
	inputFlags :=  flag.NewFlagSet("CommandFlags", flag.ContinueOnError)

	var (
		paths string
		recursive bool
	)
	inputFlags.StringVar(&paths, "paths","",  "List of paths separated by semicolons.")
	inputFlags.StringVar(&paths, "p","",  "List of paths separated by semicolons.")
	inputFlags.BoolVar(&recursive, "recursive", true, "By default, the watcher is recursive. Use --no-recursive to make the watcher non-recursive.")
	inputFlags.BoolVar(&recursive, "r", true, "By default, the watcher is recursive. Use --no-recursive to make the watcher non-recursive.")

	err = inputFlags.Parse(commandLineArgs)

	// if parsing was successful, set the CommandLineSettings values
	if err == nil {
		settings = CommandLineSettings{FolderPaths: trimArray(parsePathArray(paths)), Recursive: recursive}
	}

	return
}

func parsePathArray(pathsString string) (paths []string) {
	paths = strings.Split(pathsString, ";")
	return
}

func trimArray(inputArray []string) (outputArray []string) {
	for _, stringValue := range inputArray {
		trimmedValue := strings.Trim(stringValue, " ")
		if len(trimmedValue) > 0 {
			outputArray = append(outputArray, trimmedValue)
		}
	}
	return
}
