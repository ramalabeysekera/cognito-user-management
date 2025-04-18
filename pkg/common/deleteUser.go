package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func DeleteUser(awsConfig aws.Config, userPoolId string, userName string) error {
	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(awsConfig)

	// Prepare user deletion input parameters
	userInput := cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: &userPoolId,
		Username:   &userName,
	}

	// Set timeout context for AWS API call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete user in Cognito
	_, err := cogClient.AdminDeleteUser(ctx, &userInput)

	if err != nil {
		return err
	}

	return nil
}
