package project

import (
	"context"
	"fmt"
	"os"
	path "path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"

	"github.com/fatih/color"
	"github.com/luenci/golang-project-template/generator/template"
	"github.com/luenci/golang-project-template/internal/base"
)

type Project struct {
	Name string
	Path string
}

// NewProject new a project template.
func (p *Project) NewProject(ctx context.Context, dir string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}
	fmt.Printf("🚀 Creating project %s, please wait a moment.\n\n", p.Name)
	// todo: create project
	for _, file := range template.ParseTemplate(template.GinDemo) {
		if strings.Contains(file, ".") {
			if err := base.Createfile(to, file, ""); err != nil {
				return err
			}
		} else {
			if err := base.CreateDir(to, file); err != nil {
				return err
			}
		}
	}

	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println("			🤝 Thanks for using golang-project-template")
	fmt.Println("	📚 Tutorial: https://go-kratos.dev/docs/getting-started/start")
	return nil
}
