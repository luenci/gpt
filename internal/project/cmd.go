package project

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	timeout string
	temple  string
)

type SurveyInfo struct {
	Name      string `survey:"name"`
	GoMod     string `survey:"goMod"`
	GoVersion string `survey:"goVersion"`
}

func init() {
	// 设置默认值  默认为go
	timeout = "180s"
	temple = "gin"

	CreateCmd.Flags().StringVarP(&timeout, "timeout", "t", timeout, "time out")
	CreateCmd.Flags().StringVarP(&temple, "temple", "p", temple, "project temple")
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project template",
	Long:  `This is cmd to create a new project template.`,
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
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

	var surveyInfo SurveyInfo
	prompt := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is project name ?",
				Help:    "Created project name.",
			},
			Validate: survey.Required,
		},
		{
			Name: "goVersion",
			Prompt: &survey.Input{
				Message: "What is go version ?",
				Help:    "Created go version .",
			},
			Validate: survey.Required,
		},
		{
			Name: "goMod",
			Prompt: &survey.Input{
				Message: "What is go.mod name ?",
				Help:    "Created go.mod name.",
			},
		},
	}
	if err = survey.Ask(prompt, &surveyInfo); err != nil {
		panic(err)
	}

	p := &Project{Name: filepath.Base(surveyInfo.Name), Path: surveyInfo.Name, ProjectType: temple}
	done := make(chan error, 1)

	go func() {
		done <- p.NewProject(ctx, wd, surveyInfo.GoVersion, surveyInfo.GoMod)
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
			fmt.Fprintf(
				os.Stderr,
				"\033[31mERROR: Failed to create project(%s)\033[m\n",
				err.Error(),
			)
		}
	}
}
