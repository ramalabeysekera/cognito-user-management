// Package selections provides user interface selection functionality
package selections

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/manifoldco/promptui"
)

// SelectUserPool presents an interactive prompt to select a Cognito User Pool
// Takes an AWS config and returns the selected User Pool ID as a string
func SelectUserPool(cfg aws.Config) (string) {

	userPools, err := getAllPools(cfg)

	if err != nil {
		log.Print(err)
		return ""
	}

	if userPools == nil {
		log.Print("No User Pools found")
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
		Size: 10, // Number of items to display at once
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

func getAllPools(cfg aws.Config) ([]string, error){
    // Create slice to store User Pool IDs
    var userPools []string
    
    // Initialize Cognito client with provided config
    cogClient := cognitoidentityprovider.NewFromConfig(cfg)
    
    // Set maximum number of User Pools to retrieve
    var maxResults int32 = 20
    
    // Configure input parameters for ListUserPools API call
    listUserPoolInputs := &cognitoidentityprovider.ListUserPoolsInput{
        MaxResults: &maxResults,
    }
    
    gotAllPools := false
    for !gotAllPools {
        // Create background context
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        
        // Call ListUserPools API and handle any errors
        listUserPoolOutput, err := cogClient.ListUserPools(ctx, listUserPoolInputs)
        
        if err != nil {
            return nil, err
        }
        
        // Extract User Pool IDs from response
        for i := range listUserPoolOutput.UserPools {
            userPools = append(userPools, *listUserPoolOutput.UserPools[i].Id)
        }
        
        if listUserPoolOutput.NextToken == nil {
            gotAllPools = true
        } else {
            // Set the next token for the next iteration
            listUserPoolInputs.NextToken = listUserPoolOutput.NextToken
        }
    }
    return userPools, nil
}