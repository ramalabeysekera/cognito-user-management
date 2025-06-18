/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ramalabeysekera/cognito-user-management/config"
	"github.com/ramalabeysekera/cognito-user-management/pkg/common"
	"github.com/ramalabeysekera/cognito-user-management/pkg/helpers"
	"github.com/spf13/cobra"
)

// addtogroupCmd represents the addtogroup command
var addtogroupCmd = &cobra.Command{
	Use:   "addtogroups",
	Short: "Add a user to one or more groups in a Cognito User Pool",
	Long: `The "addtogroups" command allows you to select a user from a Cognito User Pool 
and add them to one or more groups interactively. This command simplifies group 
management by providing an intuitive CLI interface for selecting users and groups.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get selected user pool from available pools
		userPools, err := common.GetAllPools(config.AwsConfig)
		if err != nil {
			log.Println("Error fetching user pools:", err)
			return
		}
		userPool := helpers.CallSingleSelect(userPools)
		if userPool != "" {
			users, err := common.GetUsersFromPool(userPool, config.AwsConfig)
			if err != nil {
				log.Println("Error fetching users:", err)
				return
			}
			if len(users) == 0 {
				log.Println("No users found in the selected user pool.")
				return
			}
			// Let user select a user
			// Use the interactive selection function to let the user choose a user
			// This function will display the list of users and allow the user to select one
			user := helpers.CallSingleSelect(users)

			groups, err := common.GetGroupsFromPool(userPool, config.AwsConfig)
			if err != nil {
				log.Println("Error fetching groups:", err)
				return
			}
			if len(groups) == 0 {
				log.Println("No groups found in the selected user pool.")
				return
			}
			fmt.Println("Select groups to add the user to:")
			// Let user select a group
			selectedGroups := helpers.CallMultiSelect(groups)
			if err != nil {
				log.Println("Error selecting group:", err)
				return
			}

			if len(selectedGroups) == 0 {
				log.Println("No groups selected.")
				return
			}
			for _, group := range selectedGroups {
				err = common.AddUserToGroup(userPool, user, group, config.AwsConfig)
				if err != nil {
					log.Println("Error adding user to group:", err)
					return
				}
				helpers.PrintSuccessLog(fmt.Sprintf("User %s added to group %s\n", user, group))
			}
		} else {
			helpers.PrintFatalErrorLog("No user pool selected")
		}

	},
}

func init() {
	rootCmd.AddCommand(addtogroupCmd)
}
