// Package selections provides user interface selection functionality
package selections

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/manifoldco/promptui"
)

// SelectUserPool presents an interactive prompt to select a Cognito User Pool
// Takes an AWS config and returns the selected User Pool ID as a string
func SelectUserPool(cfg aws.Config) (string) {

	// Initialize Cognito client with provided config
	cogClient := cognitoidentityprovider.NewFromConfig(cfg)

	// Create background context
	ctx := context.Background()

	// Set maximum number of User Pools to retrieve
	var maxResults int32 = 20

	// Configure input parameters for ListUserPools API call
	listUserPoolInputs := cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: &maxResults,
	}

	// Call ListUserPools API and handle any errors
	listUserPoolOutput, err := cogClient.ListUserPools(ctx, &listUserPoolInputs)

	if err != nil {
		log.Print(err)
	}

	// Create slice to store User Pool IDs
	var userPools []string

	// Extract User Pool IDs from response
	for i := range listUserPoolOutput.UserPools {
		userPools = append(userPools, *listUserPoolOutput.UserPools[i].Id)
	}

	// Return empty string if 1 or fewer pools found
	if (len(listUserPoolOutput.UserPools)) <= 1 {
		return ""
	}
	
	// Configure display templates for the selection prompt
	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold | cyan }}",
		Active:   "\U0001F336 {{ . | cyan | bold }}",      // Hot pepper emoji
		Inactive: "  {{ . | white }}",
		Selected: "\U0001F525 You chose: {{ . | green }}", // Fire emoji
	}

	// Configure the selection prompt
	prompt := promptui.Select{
		Label: "Select the User Pool: ",
		Items: userPools,
		Templates: templates,
		// Define search function for filtering results
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(userPools[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
		StartInSearchMode: true,
	}

	// Run the prompt and get selected User Pool ID
	_, userPoolId, err := prompt.Run()

	if err != nil {
		log.Print(err)
	}

	// Return the selected User Pool ID
	return userPoolId
}
