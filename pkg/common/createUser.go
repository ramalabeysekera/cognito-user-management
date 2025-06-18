package common

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/ramalabeysekera/cognito-user-management/pkg/helpers"
)

func CreateUser(userPoolId string, userName string, tempPassword string, permpass bool, AwsConfig aws.Config) error {

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	// Prepare user creation input parameters
	userInput := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &userPoolId,
		Username:          &userName,
		MessageAction:     types.MessageActionTypeSuppress,
		TemporaryPassword: &tempPassword,
	}

	// Set timeout context for AWS API call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create user in Cognito
	AdminCreateUserOutput, err := cogClient.AdminCreateUser(ctx, &userInput)

	// Handle user creation response
	if err != nil {
		return err
	} else {
		if permpass {
			helpers.PrintSuccessLog("User created successfully with permanent password")
		} else {
			helpers.PrintSuccessLog("User created successfully with temporary password")
			log.Printf("Username: %s, UserStatus: %s",
				*AdminCreateUserOutput.User.Username,
				AdminCreateUserOutput.User.UserStatus)
		}
	}
	return nil
}
