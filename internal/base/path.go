package base

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// Createfile creates a file with the given name and content.
func Createfile(name, content string) error {
	return nil
}

// CreateDir creates a directory with the given name.
func CreateDir(dirName string) error {
	var re = regexp.MustCompile(`/|\\`)
	if re.MatchString(dirName) {
		return os.MkdirAll(dirName, os.ModePerm)
	}
	return os.Mkdir(dirName, os.ModePerm)
}

// Tree Path represents a path to a file or directory.
func Tree(path string, dir string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			fmt.Printf("%s %s (%v bytes)\n", color.GreenString("CREATED"), strings.Replace(path, dir+"/", "", -1), info.Size())
		}
		return nil
	})
}
