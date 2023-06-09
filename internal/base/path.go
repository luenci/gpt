package base

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/fatih/color"
)

// CreateFile creates a file with the given name and content.
func CreateFile(path, fileName, fileContent string) error {
	f, err := os.Create(filepath.Join(path, fileName))

	defer f.Close() //nolint:staticcheck

	if err != nil {
		return err
	}

	if _, err = f.Write([]byte(fileContent)); err != nil {
		return err
	}

	return nil
}

// CreateDir creates a directory with the given name.
func CreateDir(path, dirName string) error {
	re := regexp.MustCompile(`/|\\`)
	if re.MatchString(dirName) {
		return os.MkdirAll(filepath.Join(path, dirName), os.ModePerm)
	}
	return os.Mkdir(filepath.Join(path, dirName), os.ModePerm)
}

// CreateTemplateFile creates a file with the given template.
func CreateTemplateFile(path, fileName, fileTemplate string) error {
	f, err := os.Create(filepath.Join(path, fileName))
	defer f.Close() //nolint:staticcheck
	if err != nil {
		return err
	}

	t, err := template.New(fileName).Parse(fileTemplate)
	if err != nil {
		return err
	}

	err = t.Execute(f, nil)
	if err != nil {
		return err
	}

	return nil
}

// CreateArgsTemplate creates a file with arguments template.
func CreateArgsTemplate(path, fileName, ArgsTemplate string) error {
	projectName := fileName

	f, err := os.Create(filepath.Join(path, fileName))
	defer f.Close() //nolint:staticcheck
	if err != nil {
		return err
	}

	t, err := template.New(fileName).Parse(ArgsTemplate)
	if err != nil {
		return err
	}

	err = t.Execute(f, projectName)
	if err != nil {
		return err
	}

	return nil
}

// Tree Path represents a path to a file or directory.
func Tree(path, dir string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil {
			fmt.Printf(
				"%s %s (%v bytes)\n",
				color.GreenString("CREATED"),
				strings.Replace(path, dir+"/", "", -1),
				info.Size(),
			)
		}
		return nil
	})
}
