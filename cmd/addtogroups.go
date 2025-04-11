/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/ramalabeysekera/cognitousermanagement/config"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/common"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/selections"
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
		userPool := selections.SelectUserPool(config.AwsConfig)
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
			user, err := helpers.InteractiveSelection(users, "Please enter the username of the user you would like to add to a group: ")
			if err != nil {
				log.Println("Error selecting user:", err)
				return
			}
			groups, err := common.GetGroupsFromPool(userPool, config.AwsConfig)
			if err != nil {
				log.Println("Error fetching groups:", err)
				return
			}
			if len(groups) == 0 {
				log.Println("No groups found in the selected user pool.")
				return
			}
			// Let user select a group
			selectedGroups, err := helpers.InteractiveMultiSelect("Please select the groups you would like to add the user to: ", groups)
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
				greenColor := "\033[32m"
				resetColor := "\033[0m"
				log.Printf(greenColor+"User %s added to group %s\n", user, group+resetColor)
			}
		} else {
			log.Fatal("No user pool selected")
		}

	},
}

func init() {
	rootCmd.AddCommand(addtogroupCmd)
}
