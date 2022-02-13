package project

import (
	"context"
	"errors"
	"fmt"
	"os"
	path "path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	timeout     string
	projectType string
)

func init() {
	// 设置默认值  默认为go
	timeout = "180s"
	projectType = "gin"

	CreateCmd.Flags().StringVarP(&timeout, "timeout", "t", timeout, "time out")
	CreateCmd.Flags().StringVarP(&projectType, "projectType", "p", projectType, "project type")
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project template",
	Long:  `This is cmd to create a new project template.`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	name := ""
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name ?",
			Help:    "Created project name.",
		}
		err = survey.AskOne(prompt, &name)
		if err != nil || name == "" {
			return
		}
	} else {
		name = args[0]
	}
	p := &Project{Name: path.Base(name), Path: name, ProjectType: projectType}
	done := make(chan error, 1)
	go func() {
		done <- p.NewProject(ctx, wd)
	}()
	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Fprint(os.Stderr, "\033[31mERROR: project creation timed out\033[m\n")
		} else {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: failed to create project(%s)\033[m\n", ctx.Err().Error())
		}
	case err = <-done:
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: Failed to create project(%s)\033[m\n", err.Error())
		}
	}
}
