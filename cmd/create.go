/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var name, description, template string
var gitInit bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for inputs if not provided via flags
		if name == "" {
			fmt.Print("Enter the project name: ")
			name = getInput()
		}

		if description == "" {
			fmt.Print("Enter the project description: ")
			description = getInput()
		}

		if template == "" {
			template = selectTemplate()
		}

		// Display collected information
		fmt.Printf("\nProject Name: %s, Description: %s, Git Repo: %v, Template: %s\n", name, description, gitInit, template)

		// Initialize Git repository if the flag is set
		if gitInit {
			fmt.Println("Initializing Git repository...")
			err := initGitRepo(name)
			if err != nil {
				fmt.Println("Error initializing Git repository:", err)
			}
		}

		// Clone the selected template repository
		cloneRepo(template)
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the project")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the project")
	createCmd.Flags().BoolVarP(&gitInit, "git", "g", false, "Initialize a git repository")
	createCmd.Flags().StringVarP(&template, "template", "t", "", "Template for the project (web2, web3, ai)")

	rootCmd.AddCommand(createCmd)
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Function to select a template using up/down keys
func selectTemplate() string {
	templates := []string{"web2", "web3", "ai"}

	prompt := promptui.Select{
		Label: "Select a template",
		Items: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

// Clone the repository based on the template
func cloneRepo(template string) {
	var repoURL string
	switch template {
	case "web2":
		repoURL = "https://github.com/Shivam-kum-mhta/NFTMinter.git"
	case "web3":
		repoURL = "https://github.com/Shivam-kum-mhta/NFTMinter.git"
	case "ai":
		repoURL = "https://github.com/Shivam-kum-mhta/NFTMinter.git"
	default:
		fmt.Println("Invalid template selected!")
		return
	}

	fmt.Printf("Cloning repository for %s template...\n", template)
	cmd := exec.Command("git", "clone", repoURL, name) // Clone into the project directory
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error cloning repository:", err)
	} else {
		fmt.Println("Repository cloned successfully!")
	}
}

// Initialize a Git repository in the project directory
func initGitRepo(projectName string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = projectName // Set the working directory to the project folder
	return cmd.Run()
}
