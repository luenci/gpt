package project

import (
	"context"
	"fmt"
	"os"
	path "path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/luenci/gpt/generator"

	"github.com/fatih/color"
	"github.com/luenci/gpt/generator/template"
	"github.com/luenci/gpt/internal/base"
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
	// 生成空文件和目录
	for _, file := range template.ParseTemplate(template.GinDemo) {
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
	// todo: 配置参数选择是否需要默认模版
	// 生成默认模版文件
	for fileName, temlp := range generator.TemplateMap {
		if err := base.CreateTemplateFile(to, fileName, temlp); err != nil {
			return err
		}
	}
	// 生成参数模版文件
	for fileName, argsTemlp := range generator.ArgsTemplateMap {
		if err := base.CreateArgsTemplate(to, fileName, argsTemlp); err != nil {
			return err
		}
	}

	// 输出项目信息
	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println("			🤝 Thanks for using golang-project-template")
	fmt.Println("	📚 Tutorial: https://github.com/luenci/golang-project-template#readme")
	return nil
}
