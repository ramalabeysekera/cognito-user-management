package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// AdminGetUser retrieves user details from Cognito user pool
// Parameters:
//   - username: The username of the user to retrieve
//   - userPoolId: The ID of the Cognito user pool
//   - AwsConfig: AWS configuration object
//   - ctx: Context for the API call
// Returns:
//   - cognitoidentityprovider.AdminGetUserOutput: User details if successful
//   - error: Error if the operation fails
func AdminGetUser(username string, userPoolId string, AwsConfig aws.Config, ctx context.Context) (cognitoidentityprovider.AdminGetUserOutput,error) {
	
	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	// Prepare input parameters for AdminGetUser API call
	AdminGetUserInput := cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: &userPoolId,
		Username: &username,
	}

	// Call AdminGetUser API to retrieve user details
	AdminGetUserOutput , err := cogClient.AdminGetUser(ctx, &AdminGetUserInput)

	// Return empty output and error if the API call fails
	if err != nil {
		return cognitoidentityprovider.AdminGetUserOutput{}, err
	}

	// Return the user details if successful
	return *AdminGetUserOutput, nil
}
