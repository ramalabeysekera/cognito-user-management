package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// DescribeUserSignInAttr retrieves the sign-in attributes configured for a Cognito user pool
// Parameters:
//   - userPoolId: ID of the Cognito user pool to describe
//   - AwsConfig: AWS configuration for the Cognito client
//   - ctx: Context for the API call
// Returns:
//   - []string: List of configured sign-in attributes
//   - error: Any error that occurred during the operation
func DescribeUserSignInAttr(userPoolId *string, AwsConfig aws.Config, ctx context.Context) ([]string, error) {

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	// Prepare input parameters for DescribeUserPool API call
	DescribeUserPoolInput := cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: userPoolId,
	}

	// Set timeout of 10 seconds for the API call
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Call Cognito API to get user pool details
	DescribeUserPoolOutput, err := cogClient.DescribeUserPool(ctx, &DescribeUserPoolInput)

	// Return error if API call fails
	if err != nil {
		return nil, err
	}

	// Initialize slice to store attributes
	var attrs []string

	// Get sign-in identifiers (username attributes) from the user pool
	signInOptions := DescribeUserPoolOutput.UserPool.UsernameAttributes
	// Convert each attribute to string and append to result slice
	for _, attr := range signInOptions {
		attrs = append(attrs, string(attr))
	}

	// Return the list of sign-in attributes
	return attrs, nil
}
