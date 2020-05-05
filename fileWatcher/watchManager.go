package fileWatcher

import (
	"errors"
	"github.com/radovskyb/watcher"
	"strings"
)

type WatchManager struct {
	Watcher    *watcher.Watcher
	FolderList map[string]Folder
}

func NewWatchManager() *WatchManager {
	wm := new(WatchManager)
	wm.Watcher = watcher.New()
	wm.FolderList = make(map[string]Folder)
	return wm
}

func (wm *WatchManager) AddFolder(path string, recursive bool) (err error) {
	path = strings.Trim(path, " ")
	if len(path) == 0 {
		return errors.New("path is empty")
	}
	newFolder := &Folder{Path: path, Recursive: recursive}
	if recursive {
		wm.Watcher.AddRecursive(path)
	} else {
		wm.Watcher.Add(path)
	}
	wm.FolderList[path] = *newFolder
	return
}

// End the watching for all folders
func (wm *WatchManager) Close() {
	wm.Watcher.Close()
}

func (wm *WatchManager) RemoveFolder(path string) (err error) {
	path = strings.Trim(path, " ")
	if len(path) == 0 {
		return errors.New("path is empty")
	}
	err = wm.Watcher.Remove(path)
	delete(wm.FolderList, path)
	return
}
