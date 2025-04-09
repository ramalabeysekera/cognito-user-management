package config

import (
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
    "os"
)

var AwsConfig aws.Config

func init() {
    if len(os.Args) > 1 {
        AwsConfig = helpers.LoadAwsConfig()
    }
}