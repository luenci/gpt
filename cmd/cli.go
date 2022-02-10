package cmd

import (
	"fmt"
	"os"

	"github.com/luenci/golang-project-template/internal/project"
	"github.com/luenci/golang-project-template/internal/version"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gpt",
		Short: "A simple command create golang projects templates",
		Long: `This command creates golang projects templates,
			default templates is gin-demo( https://github.com/Lucareful/gin-demo)`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd")
		},
	}
)

func init() {
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(project.CreateCmd)
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
