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
func SetPermanentPassword(setPermanentPasswordInput SetPermanentPasswordInput, AwsConfig aws.Config, ctx context.Context) error{

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Prepare input for AdminSetUserPassword API call
	adminSetPasswordInput := cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: &setPermanentPasswordInput.UserPoolId,
		Username:   &setPermanentPasswordInput.Username,
		Password:   &setPermanentPasswordInput.Password,
		Permanent:  true, 
	}

	// Call AdminSetUserPassword API
	_ , err := cogClient.AdminSetUserPassword(ctx, &adminSetPasswordInput)

	if err != nil {
		return err
	} 

	return nil
}