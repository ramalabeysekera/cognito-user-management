package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type SetPermanentPasswordInput struct {
	UserPoolId string
	Username string
	Password string
}

// setPermanentPassword handles the password setting logic for a Cognito user
func SetPermanentPassword(userPoolId string, username string, password string, AwsConfig aws.Config, ctx context.Context) (cognitoidentityprovider.AdminSetUserPasswordOutput, error){

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Prepare input for AdminSetUserPassword API call
	adminSetPasswordInput := cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: &userPoolId,
		Username:   &username,
		Password:   &password,
		Permanent:  true, 
	}

	// Call AdminSetUserPassword API
	AdminSetUserPasswordOutput , err := cogClient.AdminSetUserPassword(ctx, &adminSetPasswordInput)

	if err != nil {
		return cognitoidentityprovider.AdminSetUserPasswordOutput{}, err
	} 

	return *AdminSetUserPasswordOutput, nil
}