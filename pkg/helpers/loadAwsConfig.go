// Package helpers provides utility functions for AWS configuration and profile management
package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// LoadAwsConfig loads and returns an AWS configuration based on the selected profile.
// It performs the following:
// - Gets an AWS profile name from user selection
// - Loads the AWS config for that profile
// - Validates the credentials by making an STS GetCallerIdentity call
// - Prints account and identity information
// Returns the AWS configuration or exits on error
func LoadAwsConfig() aws.Config {

	// Get AWS profile name from user selection
	profile := SelectAwsProfile()

	// Validate that a profile was selected
	if profile == "" {
		log.Fatal("Need a profile to continue")
	}

	// Load AWS configuration from default config sources
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	// Handle any config loading errors
	if err != nil {
		panic(err)
	}

	// Print the profile being used
	fmt.Println("Using profile:", profile)

	// Print the configured AWS region
	fmt.Println("Using region: ", cfg.Region)
	
	// Create a new STS client using the loaded config
	client := sts.NewFromConfig(cfg)

	// Get the caller identity to validate credentials and get account info
	result, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		panic("failed to get caller identity, " + err.Error())
	}

	// Print account and identity information
	fmt.Println("Using account ID:", *result.Account)
	fmt.Println("Using caller identity:", *result.Arn)

	return cfg
}
