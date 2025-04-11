package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func GetGroupsFromPool(userPoolId string, awsConfig aws.Config) ([]string, error) {
	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(awsConfig)

	// Create the input for the ListGroups API call
	input := &cognitoidentityprovider.ListGroupsInput{
		UserPoolId: &userPoolId,
	}

	// Call the ListGroups API
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	output, err := cogClient.ListGroups(ctx, input)
	if err != nil {
		return nil, err
	}

	// Extract group names from the output
	var groups []string
	for _, group := range output.Groups {
		groups = append(groups, *group.GroupName)
	}

	return groups, nil
}