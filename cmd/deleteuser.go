/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ramalabeysekera/cognitousermanagement/config"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/common"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
	"github.com/spf13/cobra"
)

// deleteuserCmd represents the deleteuser command
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "Delete a user from an AWS Cognito user pool",
	Long: `Delete a user from an AWS Cognito user pool.
	
This command will:
1. Let you select a user pool from your AWS account
2. Display all users in that pool
3. Allow you to select a user to delete
4. Ask for confirmation before deletion
5. Delete the selected user from the pool

Example:
  cognitousermanagement deleteuser`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get user pool selection from user
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
			fmt.Println("Select a user to delete:")
			// Display interactive user selection prompt
			user := helpers.CallSingleSelect(users)

			// Confirm deletion with user
			fmt.Print("Are you sure you want to delete this user? (y/n): ")
			reader := bufio.NewReader(os.Stdin)
			confirmation, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading confirmation:", err)
				return
			}
			confirmation = strings.TrimSpace(confirmation)
			confirmation = strings.ToLower(confirmation)
			if confirmation != "y" {
				log.Println("User deletion cancelled.")
				return
			}

			// Delete the user from Cognito
			err = common.DeleteUser(config.AwsConfig, userPool, user)

			if err != nil {
				log.Println("Error deleting user:", err)
				return
			} else {
				helpers.PrintSuccessLog(fmt.Sprintf("User %s deleted successfully from pool %s\n", user, userPool))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)
}
