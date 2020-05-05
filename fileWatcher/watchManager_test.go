package fileWatcher

import (
	"testing"
)

func TestWatchManager_AddRemoveFolder(t *testing.T) {
	watchMan := NewWatchManager()
	var err error

	// try to add folders
	err = watchMan.AddFolder(".", true)
	if len(watchMan.FolderList) != 1 || err != nil {
		t.Error("Folder was not added")
	}

	// remove one folder
	err = watchMan.RemoveFolder(".")
	if len(watchMan.FolderList) != 0 || err != nil {
		t.Error("Folder was not removed")
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
		{name: "a path that does not exist", path: "/home/path", recursive: true, wantErr: true},
		{name: "Add a path recursive", path: "./", recursive: true, wantErr: false},
		{name: "Add a path non-recursive", path: ".", recursive: true, wantErr: false},
	}
	for _, tt := range tests {
		watchMan := NewWatchManager()
		t.Run(tt.name, func(t *testing.T) {

			if err := watchMan.AddFolder(tt.path, true); (err != nil) != tt.wantErr {
				t.Errorf("AddFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
			// if this scenario should generate an error, skip the check for the recursive flag
			if !tt.wantErr && watchMan.FolderList[tt.path].Recursive != tt.recursive {
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
