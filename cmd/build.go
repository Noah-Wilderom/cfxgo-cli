/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Noah-Wilderom/cfxgo-cli/types"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		configFile := path.Join(cwd, "cfx-go.config.yaml")
		if _, err := os.Stat(configFile); err != nil {
			panic(err)
		}

		configContent, err := os.ReadFile(configFile)
		if err != nil {
			panic(err)
		}

		var config types.Config
		if err := yaml.Unmarshal([]byte(configContent), &config); err != nil {
			panic(err)
		}

		buildDir := path.Join(cwd, "build")
		cleanBuildDir(buildDir)

		if err := buildGo(cwd, buildDir, &config); err != nil {
			log.Fatal(err)
		}

		if err := buildTypescript(cwd, buildDir, &config); err != nil {
			log.Fatal(err)
		}
	},
}

func buildGo(cwd string, buildDir string, config *types.Config) error {
	if config.Server.Go.Enable {
		// if err := os.MkdirAll(path.Join(buildDir, "server"), os.ModePerm); err != nil {
		// 	panic(err)

		// }

		// if err := os.Setenv("GOOS", "js"); err != nil {
		// 	return err
		// }
		// if err := os.Setenv("GOARCH", "wasm"); err != nil {
		// 	return err
		// }

		// buildCmd := exec.Command("go", "build", "-o", path.Join(buildDir, "server", "go.wasm"), path.Join(cwd, "src", "server"))
		buildCmd := exec.Command("docker", "compose", "up", "-d", "--build")
		buildCmd.Dir = cwd
		if err := buildCmd.Run(); err != nil {
			return err
		}
	}

	if config.Client.Go.Enable {
		if err := os.MkdirAll(path.Join(buildDir, "client"), os.ModePerm); err != nil {
			panic(err)
		}

		buildCmd := exec.Command(strings.Split(config.Client.Go.Exec, " ")[0], strings.Split(config.Client.Go.Exec, " ")[1:]...)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func buildTypescript(cwd string, buildDir string, config *types.Config) error {
	installPackages(cwd)

	esbuildConfig := path.Join(cwd, "esbuild.config.mjs")
	if err := os.MkdirAll(path.Join(buildDir, "client"), os.ModePerm); err != nil {
		panic(err)
	}

	if config.Server.Typescript.Enable {
		cmd := exec.Command("node", esbuildConfig, config.Server.Typescript.Main, path.Join(buildDir, "server"))
		cmd.Dir = cwd
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if config.Client.Typescript.Enable {
		cmd := exec.Command("node", esbuildConfig, config.Client.Typescript.Main, path.Join(buildDir, "client"))
		cmd.Dir = cwd
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if config.Shared.Typescript.Enable {
		cmd := exec.Command("node", esbuildConfig, config.Client.Typescript.Main, path.Join(buildDir, "shared"))
		cmd.Dir = cwd
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func installPackages(cwd string) {
	installCmd := exec.Command("npm", "install", "typescript", "esbuild", "esbuild-obfuscator-plugin", "--save-dev")
	installCmd.Dir = cwd
	if err := installCmd.Run(); err != nil {
		panic(err)
	}
}

func cleanBuildDir(buildDir string) {
	_ = os.RemoveAll(buildDir)
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
