/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ramalabeysekera/cognitousermanagement/config"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/common"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
	"github.com/spf13/cobra"
)

// setpasswordCmd represents the setpassword command
var setpasswordCmd = &cobra.Command{
	Use:   "setpassword",
	Short: "Set a permanent password for a Cognito user",
	Long: `The setpassword command allows you to set a permanent password for a user in an AWS Cognito user pool.
	
It will:
1. Let you select a user pool from available pools
2. Display list of users in the selected pool
3. Allow you to select a user
4. Prompt for a new password
5. Set the permanent password for the selected user`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get selected user pool from available pools by displaying interactive selection
				// Get selected user pool from available pools
				userPools, err := common.GetAllPools(config.AwsConfig)
				if err != nil {
					log.Println("Error fetching user pools:", err)
					return
				}
				userPool := helpers.CallSingleSelect(userPools)
		if userPool != "" {
			// Fetch all users from the selected pool
			users, err := common.GetUsersFromPool(userPool, config.AwsConfig)
			if err != nil {
				log.Println("Error fetching users:", err)
				return
			}
			// Validate that users exist in the pool
			if len(users) == 0 {
				log.Println("No users found in the selected user pool.")
				return
			}
			fmt.Println("Select a user to set a permanent password:")
			// Display interactive user selection prompt
      
			user := helpers.CallSingleSelect(users)


			// Get the password from the user via stdin
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Please enter the new password: ")
			password, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading password:", err)
				return
			}

			// Remove the newline character from the password input
			password = strings.TrimSpace(password)

			// Create a context for the operation
			ctx := context.Background()

			// Call AWS Cognito API to set the permanent password for the user
			_, err = common.SetPermanentPassword(userPool, user, password, config.AwsConfig, ctx)

			if err != nil {
				log.Println("Error setting password:", err)
				return
			}
			helpers.PrintSuccessLog(fmt.Sprintf("Password set successfully for user %s\n", user))
		}
	},
}

func init() {
	rootCmd.AddCommand(setpasswordCmd)
}
