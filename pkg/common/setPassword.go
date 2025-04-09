package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// SetPermanentPasswordInput defines the input parameters required to set a permanent password
// for a Cognito user. It includes the user pool ID, username and the new password.
type SetPermanentPasswordInput struct {
	UserPoolId string // The ID of the Cognito user pool
	Username string   // The username of the Cognito user
	Password string   // The new password to be set
}

// SetPermanentPassword handles the password setting logic for a Cognito user
// Parameters:
//   - userPoolId: The ID of the Cognito user pool
//   - username: The username of the Cognito user
//   - password: The new password to be set
//   - AwsConfig: AWS configuration object
//   - ctx: Context for the operation
// Returns:
//   - AdminSetUserPasswordOutput: Response from Cognito API
//   - error: Any error that occurred during the operation
func SetPermanentPassword(userPoolId string, username string, password string, AwsConfig aws.Config, ctx context.Context) (cognitoidentityprovider.AdminSetUserPasswordOutput, error){

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	// Create context with timeout of 10 seconds to prevent hanging operations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure resources are cleaned up when function returns

	// Prepare input for AdminSetUserPassword API call
	// Setting Permanent to true makes this a permanent password change
	adminSetPasswordInput := cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: &userPoolId,
		Username:   &username,
		Password:   &password,
		Permanent:  true, 
	}

	// Call AdminSetUserPassword API to set the new password
	AdminSetUserPasswordOutput , err := cogClient.AdminSetUserPassword(ctx, &adminSetPasswordInput)

	// If there's an error, return empty output and the error
	if err != nil {
		return cognitoidentityprovider.AdminSetUserPasswordOutput{}, err
	} 

	// Return the API response if successful
	return *AdminSetUserPasswordOutput, nil
}