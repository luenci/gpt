package project

import (
	"context"
	"fmt"
	"os"
	path "path/filepath"
	"strings"

	"github.com/luenci/gpt/generator/content"

	"github.com/AlecAivazis/survey/v2"
	"github.com/luenci/gpt/generator"
	"github.com/luenci/gpt/generator/template"

	"github.com/fatih/color"
	"github.com/luenci/gpt/internal/base"
)

type Project struct {
	_           struct{}
	Name        string
	Path        string
	ProjectType string
}

// NewProject new a project template.
func (p *Project) NewProject(ctx context.Context, dir string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("ð« %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "ð Do you want to override the folder ?",
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
	fmt.Printf("ð Creating project %s, please wait a moment.\n\n", p.Name)

	// æ ¹æ®ä¼ å¥çé¡¹ç®ç±»åçæç©ºæä»¶åç®å½
	switch p.ProjectType {
	case "gin":
		for _, file := range template.ParseTemplate(template.GinWeb) {
			if strings.Contains(file, ".") {
				if err := base.CreateFile(to, file, ""); err != nil {
					return err
				}
			} else {
				if err := base.CreateDir(to, file); err != nil {
					return err
				}
			}
		}
		if err := base.CreateTemplateFile(to, "README.md", content.GinReadMe); err != nil {
			return err
		}
	case "DDD":
		for _, file := range template.ParseTemplate(template.DDD) {
			if strings.Contains(file, ".") {
				if err := base.CreateFile(to, file, ""); err != nil {
					return err
				}
			} else {
				if err := base.CreateDir(to, file); err != nil {
					return err
				}
			}
		}
		if err := base.CreateTemplateFile(to, "README.md", content.DDDReadMe); err != nil {
			return err
		}
	}

	// çæé¡¹ç®æ¨¡æ¿README

	// çæéç¨çæ¨¡çæä»¶ååå®¹
	for fileName, temlp := range generator.TemplateMap {
		if err := base.CreateTemplateFile(to, fileName, temlp); err != nil {
			return err
		}
	}

	// çæåæ°æ¨¡çæä»¶
	for fileName, argsTemlp := range generator.ArgsTemplateMap {
		if err := base.CreateArgsTemplate(to, fileName, argsTemlp); err != nil {
			return err
		}
	}

	// è¾åºé¡¹ç®ä¿¡æ¯
	base.Tree(to, dir)

	fmt.Printf("\nðº Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("ð» Use the following command to start the project ð:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println("			ð¤ Thanks for using gpt (golang-project-template)")
	fmt.Println("	ð Tutorial: https://github.com/luenci/gpt#readme")
	return nil
}
