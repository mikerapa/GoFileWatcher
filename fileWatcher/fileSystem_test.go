package fileWatcher

import "testing"

func Test_isValidDirPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{"empty path", "", false},
		{"just a slash", "/", true},
		{"fake path", "/thing", false},
		{"dot", ".", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidDirPath(tt.path); got != tt.want {
				t.Errorf("isValidDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
