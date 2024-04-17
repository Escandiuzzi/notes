/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	folder  string
	title   string
	content string
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:     "note",
	Aliases: []string{"add", "make", "n", "a"},
	Short:   "Quick way to create notes from your terminal",
	Long:    `Notes is a CLI tool that allows you to create notes from your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {

		applescriptCode := fmt.Sprintf(`
		tell application "Notes"
			set newNote to make new note at folder "%s" with properties {name:"%s", body:"%s"}
		end tell`,
			folder, title, content)

		applescriptCode = strings.TrimSpace(applescriptCode)

		osCommand := exec.Command("osascript", "-e", applescriptCode)

		output, err := osCommand.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing AppleScript:", err)
			fmt.Println("Command output:", string(output))
			return
		}

		fmt.Println("Note added successfully :)")
	},
}

func init() {
	rootCmd.AddCommand(noteCmd)

	rootCmd.PersistentFlags().StringVarP(&folder, "folder", "f", "Notes", "Defines the folder to create the note in")
	rootCmd.MarkFlagRequired("folder")
	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Defines the title of the note")
	rootCmd.PersistentFlags().StringVarP(&content, "content", "c", "Content", "Defines the content of the note")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
