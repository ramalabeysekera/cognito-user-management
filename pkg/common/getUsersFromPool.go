package common

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// GetUsersFromPool retrieves all users from a Cognito user pool
// Parameters:
//   userPoolId: ID of the Cognito user pool to query
//   awsConfig: AWS configuration object
// Returns:
//   []string: Slice containing usernames of all users in the pool
//   error: Error if the operation fails
func GetUsersFromPool(userPoolId string, awsConfig aws.Config) ([]string, error) {

	// Flag to track if all users have been retrieved
	var allUsersRetrieved bool

	// Slice to store usernames
	var users []string

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(awsConfig)

	// Create the input for the ListUsers API call
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: &userPoolId,
	}

	// Create context with timeout for the API calls
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Loop until all users are retrieved using pagination
	for !allUsersRetrieved {
	
	// Call Cognito API to get batch of users
	output, err := cogClient.ListUsers(ctx, input)
	if err != nil {
		return nil, err
	}

	// Extract usernames from the response and append to users slice
	for _, user := range output.Users {
		users = append(users, *user.Username)
	}

	// Check if there are more users to retrieve
	if output.PaginationToken == nil {
		allUsersRetrieved = true
	} else {
		// Set pagination token for next batch
		input.PaginationToken = output.PaginationToken
	}
}
	// Log the total number of users found
	fmt.Printf("%v users found in the pool %v\n", len(users), userPoolId)
	return users, nil
}
