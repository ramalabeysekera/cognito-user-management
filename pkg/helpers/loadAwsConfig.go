package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func LoadAwsConfig() aws.Config {

	profile := SelectAwsProfile()

	if profile == "" {
		log.Fatal("Need a profile to continue")
	}

	// Load AWS configuration from default config sources
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

	if err != nil {
		panic(err)
	}

	fmt.Println("Using profile:", profile)

	fmt.Println("Using region: ", cfg.Region)
	
	// Create a new STS client
	client := sts.NewFromConfig(cfg)

	// Get the caller identity
	result, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		panic("failed to get caller identity, " + err.Error())
	}

	fmt.Println("Using account ID:", *result.Account)
	fmt.Println("Using caller identity:", *result.Arn)

	return cfg
}
