package common

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func AddUserToGroup(userPoolId string, userName string, groupName string, awsConfig aws.Config) error {

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(awsConfig)

	// Create the input for the AdminAddUserToGroup API call
	input := &cognitoidentityprovider.AdminAddUserToGroupInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(userName),
		GroupName:  aws.String(groupName),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Call the AdminAddUserToGroup API
	_, err := cogClient.AdminAddUserToGroup(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to add user %s to group %s: %w", userName, groupName, err)
	}

	return nil
}