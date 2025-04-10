// Package config provides AWS configuration functionality
package config

import (
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
    "os"
)

// AwsConfig holds the AWS configuration settings
var AwsConfig aws.Config

// init initializes the AWS configuration when command line arguments are provided
func init() {
    // Only load config if command line args exist
    if len(os.Args) > 1 {
        AwsConfig = helpers.LoadAwsConfig()
    }
}
