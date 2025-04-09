package common

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func DescribeUserSignInAttr(userPoolId *string, AwsConfig aws.Config, ctx context.Context) ([]string, error) {

	// Initialize Cognito client with AWS configuration
	cogClient := cognitoidentityprovider.NewFromConfig(AwsConfig)

	DescribeUserPoolInput := cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: userPoolId,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	DescribeUserPoolOutput, err := cogClient.DescribeUserPool(ctx, &DescribeUserPoolInput)

	if err != nil{
		return nil, err
	}

	var attrs []string

	// Get sign-in identifiers (username attributes) from the user pool
	signInOptions := DescribeUserPoolOutput.UserPool.UsernameAttributes
	for _, attr := range signInOptions {
		attrs = append(attrs, string(attr))
	}


	return attrs, nil
}