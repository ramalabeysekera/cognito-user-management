/*
Copyright © 2025 Ramal Abeysekera ramal.abeysekera@hotmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const banner = `
   ____  _    _ __  __ 
  / ___|| |  | |  \/  |
 | |    | |  | | |\/| |
 | |___ | |__| | |  | |
  \____| \____/|_|  |_|

Cognito User Management
Author: Ramal Abeysekera
Version: 1.3.1
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cognitousermanagement",
	Short: "A CLI tool for managing AWS Cognito users",
	Long: `CognitoUserManagement is a command-line application designed to simplify 
the management of users in AWS Cognito user pools. 

With this tool, you can perform various operations such as creating, updating, 
deleting, and listing users, as well as managing groups and user attributes. 
It is built to streamline user management tasks for developers and administrators 
working with AWS Cognito.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Print(banner)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cognitousermanagement.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
