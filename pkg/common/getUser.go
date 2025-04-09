package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func AdminGetUser(username string, userPoolId string, AwsConfig aws.Config, ctx context.Context) (cognitoidentityprovider.AdminGetUserOutput,error) {
	
	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	AdminGetUserInput := cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &userPoolId,
		Username: &username,
	}

	// Call AdminSetUserPassword API
	AdminGetUserOutput , err := cogClient.AdminGetUser(ctx, &AdminGetUserInput)

	if err != nil {
		return cognitoidentityprovider.AdminGetUserOutput{}, err
	}

	return *AdminGetUserOutput, nil
}