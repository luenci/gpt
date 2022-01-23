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
func Createfile(path, fileName, fileContent string) error {
	f, err := os.Create(filepath.Join(path, fileName))
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(fileContent))
	return nil
}

// CreateDir creates a directory with the given name.
func CreateDir(path, dirName string) error {
	var re = regexp.MustCompile(`/|\\`)
	if re.MatchString(dirName) {
		return os.MkdirAll(filepath.Join(path, dirName), os.ModePerm)
	}
	return os.Mkdir(filepath.Join(path, dirName), os.ModePerm)
}

// Tree Path represents a path to a file or directory.
func Tree(path, dir string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			fmt.Printf("%s %s (%v bytes)\n", color.GreenString("CREATED"), strings.Replace(path, dir+"/", "", -1), info.Size())
		}
		return nil
	})
}
