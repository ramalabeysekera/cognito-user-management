package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func GetUsersFromPool(userPoolId string, awsConfig aws.Config) ([]string, error) {
	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(awsConfig)

	// Create the input for the ListUsers API call
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: &userPoolId,
	}

	// Call the ListUsers API
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	output, err := cogClient.ListUsers(ctx, input)
	if err != nil {
		return nil, err
	}

	// Extract user names from the output
	var users []string
	for _, user := range output.Users {
		users = append(users, *user.Username)
	}

	return users, nil
}
