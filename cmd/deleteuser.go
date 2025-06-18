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

	"github.com/ramalabeysekera/cognito-user-management/config"
	"github.com/ramalabeysekera/cognito-user-management/pkg/common"
	"github.com/ramalabeysekera/cognito-user-management/pkg/helpers"
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
			fmt.Println("Select a users to delete:")
			// Display interactive user selection prompt
			usersToBeDeleted := helpers.CallMultiSelect(users)

			if len(usersToBeDeleted) == 0 {
				log.Println("No users selected for deletion.")
				return
			}

			for i := range usersToBeDeleted {
				user := usersToBeDeleted[i]
				// Confirm deletion with user
				fmt.Printf("Are you sure you want to delete user %v (y/n): ", user)
				reader := bufio.NewReader(os.Stdin)
				confirmation, err := reader.ReadString('\n')
				if err != nil {
					log.Println("Error reading confirmation:", err)
					return
				}
				confirmation = strings.TrimSpace(confirmation)
				confirmation = strings.ToLower(confirmation)
				if confirmation != "y" {
					helpers.PrintWarningErrorLog(fmt.Sprintf("User %s deletion cancelled.\n", user))
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

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteuserCmd)
}
