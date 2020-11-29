# GoFileWatcher

GoFileWatcher is a command line utility for monitoring file changes. 
If you are looking for a go library for use within another application, see 
https://github.com/mikerapa/FolderWatcher. 

## Usage

Running the application in linux

`./GoFileWatcher`

Running the application in Windows

`GoFileWatcher.exe`

Command line flags

`      --help         Show context-sensitive help (also try --help-long and
                     --help-man).`
                     
  `-p, --paths=PATHS  List of paths separated by semicolons.`
  
  `-r, --recursive    By default, the watcher is recursive. Use --no-recursive to
                     make the watcher non-recursive.`
                     
### Menu System
If application starts with a folder path specified on command line, the application will start watching
the specified folder without displaying the main menu. The main menu will display if a path has not been 
provided. You can return to the main menu by pressing Enter at any time. File events are not displayed
while the menu is active.


![Main Menu](./images/mainmenu.png "Main Menu")

| Action | Description |
| ------- | -------|
| Add Folder | specify a folder path to watch |
| List Folders | see a list of folders the application is watching |
| Remove Folder | choose a folder path to stop watching |
| Resume Watch | Exit the main menu and return to viewing file events |
| Exit | Quit the application |


### Output 
// TODO need to update this image

![Main Menu](./images/listevents.png "Main Menu")

While watching folders, add, delete and update events are captured.

## Knows Limitations
1. Hidden files are not tracked. 
2. Files moved from a watched folder to an unwatched folder will be recorded as a Remove event.

## Future development 
- [ ] Hidden File - Add the ability to track events for hidden files. 

## Attributions 

The menus are built using the promptui package. 
Copyright (c) 2017, Arigato Machine Inc. All rights reserved. Seee https://github.com/manifoldco/promptui
for more details.

The Aurora package is used to produce the color output. See https://github.com/logrusorgru/aurora.

This application uses the Kinpin package for dealing with command line arguments. 
github.com/alecthomas/kingpin

