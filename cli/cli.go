package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	commands string

	rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "A simple command create golang projects templates",
		Long: `This command creates golang projects templates,
			default templates is gin-demo( https://github.com/Lucareful/gin-demo)`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cli")
			for _, command := range commands {
				fmt.Println(command)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("rootflag", "r", "two", "")
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
