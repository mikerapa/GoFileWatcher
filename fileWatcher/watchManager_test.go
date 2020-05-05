package fileWatcher

import (
	"testing"
)

func TestWatchManager_AddRemoveFolder(t *testing.T) {
	watchMan := NewWatchManager()

	// try to add folders
	watchMan.AddFolder("/base/f1", true)
	if len(watchMan.FolderList) != 1 {
		t.Error("Folder was not added")
	}

	// try to add a second folder
	watchMan.AddFolder("base/f2", false)
	if len(watchMan.FolderList) != 2 {
		t.Error("Folder was not added")
	}

	// remove one folder
	watchMan.RemoveFolder("/base/f1")
	if len(watchMan.FolderList) != 1 {
		t.Error("Folder was not added")
	}

	// Remove second folder
	watchMan.RemoveFolder("base/f2")
	if len(watchMan.FolderList) != 0 {
		t.Error("Folder was not added")
	}

}

func TestWatchManager_AddFolder(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		recursive bool
		wantErr   bool
	}{
		{name: "empty path", path: "", recursive: false, wantErr: true},
		{name: "Add a path recursive", path: "/home/path", recursive: true, wantErr: false},
		{name: "Add a path recursive", path: "/home/path2", recursive: true, wantErr: false},
	}
	for _, tt := range tests {
		watchMan := NewWatchManager()
		t.Run(tt.name, func(t *testing.T) {

			if err := watchMan.AddFolder(tt.path, true); (err != nil) != tt.wantErr {
				t.Errorf("AddFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
			if watchMan.FolderList[tt.path].Recursive != tt.recursive {
				t.Error("AddFolder() recursive flag not set correctly")
			}
		})
	}
}

func TestWatchManager_RemoveFolder(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{name: "empty path", path: "", wantErr: true},
		{name: "path that is not in the collection", path: "/home/path", wantErr: false},
	}
	for _, tt := range tests {
		watchMan := NewWatchManager()
		t.Run(tt.name, func(t *testing.T) {

			if err := watchMan.RemoveFolder(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
