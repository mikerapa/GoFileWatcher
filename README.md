# GoFileWatcher

GoFileWatcher is a command line utility for monitoring file changes. 
If you are looking for a go library for use within another application, see 
github.com/radovskyb/watcher, on which this application is based. 

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
The main menu is displayed when the application is first started, unless a path
is passed in as a flag. 

![Main Menu](./images/mainmenu.png "Main Menu")

**Add Folder:** Add a folder to the list of folders being watched

**List Folders:** Show a list of folders currently being watched

**Remove Folder:** See a list of folders currently watched and select one to remove
from the list and stop watching. 

**Resume Watch:** Exit the main menu and return to viewing file events

**Exit:** Quit the application 

### Output 
![Main Menu](./images/listevents.png "Main Menu")

While watching folders, add, delete and update events are captured.