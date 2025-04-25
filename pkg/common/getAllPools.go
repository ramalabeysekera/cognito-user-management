package common

import (
	"context"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func GetAllPools(cfg aws.Config) ([]string, error) {
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