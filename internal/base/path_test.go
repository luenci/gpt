package base

import (
	"os"
	"strings"
	"testing"
)

func TestCreateDir(t *testing.T) {
	type args struct {
		dirName string
	}
	tests := []struct {
		name     string
		args     args
		cleanAll bool
	}{
		{"linux", args{"demo"}, false},
		{"linux", args{"demo1/demo2"}, true},
		{"windows", args{"demo"}, false},
		{"windows", args{"demo1\\demo2"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDir("", tt.args.dirName); err != nil {
				t.Errorf("CreateDir() error = %v", err)
			}
			// 愿天堂没有单元测试（清理逻辑写的头疼）
			if tt.cleanAll {
				switch tt.name {
				case "linux":
					if err := os.RemoveAll(strings.Split(tt.args.dirName, "/")[0]); err != nil {
						t.Errorf("linux os.RemoveAll() error = %v", err)
					}
				case "windows":
					if err := os.RemoveAll(strings.Split(tt.args.dirName, "\\")[0]); err != nil {
						t.Errorf("windows os.RemoveAll() error = %v", err)
					}
				}
			}
			if err := os.RemoveAll(tt.args.dirName); err != nil {
				t.Errorf("RemoveAll() error = %v", err)
			}
		})
	}
}
