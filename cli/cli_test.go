package cli

import (
	"reflect"
	"strings"
	"testing"
)

func Test_getCommandLineSettings(t *testing.T) {
	tests := []struct {
		name         string
		argsString   string
		wantSettings CommandLineSettings
		wantErr      bool
	}{
		{"1 path, no Recursive", "-p /usr/mgr", CommandLineSettings{FolderPaths: []string{"/usr/mgr"}, Recursive: true}, false},
		{"1 path with spaces", "-p /usr/mgr ", CommandLineSettings{FolderPaths: []string{"/usr/mgr"}, Recursive: true}, false},
		{"1 path, Recursive false", "-p /usr/mgr --no-recursive--no-recursive", CommandLineSettings{FolderPaths: []string{"/usr/mgr"}, Recursive: false}, false},
		{"2 paths", "-p /usr/mgr;home/user1/folder2", CommandLineSettings{FolderPaths: []string{"/usr/mgr", "home/user1/folder2"}, Recursive: true}, false},
		{"0 paths", "-p", CommandLineSettings{FolderPaths: []string{""}, Recursive: true}, true},
		{"1 path semicolon at the end", "-p /usr/mgr;", CommandLineSettings{FolderPaths: []string{"/usr/mgr"}, Recursive: true}, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSettings, err := GetCommandLineSettings(strings.Split(tt.argsString, " "))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommandLineSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if getCommandLineSetting returned an error, it doesn't matter what the settings are
			if err==nil && !reflect.DeepEqual(gotSettings, tt.wantSettings) {
				t.Errorf("GetCommandLineSettings() gotSettings = %v, want %v", gotSettings, tt.wantSettings)
			}
		})
	}
}