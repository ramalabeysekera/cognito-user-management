// Package selections provides user interface selection functionality
package selections

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
)

// SelectUserPool presents an interactive prompt to select a Cognito User Pool
// Takes an AWS config and returns the selected User Pool ID as a string
func SelectUserPool(cfg aws.Config) string {
	// Get all available user pools
	userPools, err := getAllPools(cfg)

	if err != nil {
		log.Print(err)
		return ""
	}

	if len(userPools) == 0 {
		log.Print("No User Pools found")
		return ""
	}

	// Use our Bubbletea-based selection component
	userPoolId, err := helpers.InteractiveSelection(userPools, "Select the User Pool:")
	if err != nil {
		log.Print(err)
		return ""
	}

	return userPoolId
}

// getAllPools retrieves all user pools from AWS Cognito
func getAllPools(cfg aws.Config) ([]string, error) {
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
