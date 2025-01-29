/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Println("Initializing", projectName)

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		basePath := path.Join(cwd, projectName)
		if _, err := os.Stat(basePath); os.IsNotExist(err) == false {
			fmt.Println("Error:", projectName, "already exists")
			os.Exit(1)
		}

		err = os.Mkdir(basePath, 0755)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		err = cloneBoilerplate(basePath, "")
	},
}

func cloneBoilerplate(dir string, url string) error {
	cloneCmd := exec.Command("git", "clone", url, ".")
	cloneCmd.Dir = dir

	return cloneCmd.Run()
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
