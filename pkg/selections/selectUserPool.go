package selections

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/manifoldco/promptui"
)

func SelectUserPool(cfg aws.Config) (string) {

	// Initialize Cognito client
	cogClient := cognitoidentityprovider.NewFromConfig(cfg)

	ctx := context.Background()

	var maxResults int32 = 20

	listUserPoolInputs := cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: &maxResults,
	}

	listUserPoolOutput, err := cogClient.ListUserPools(ctx, &listUserPoolInputs)

	if err != nil {
		log.Print(err)
	}

	var userPools []string

	for i := range listUserPoolOutput.UserPools {
		userPools = append(userPools, *listUserPoolOutput.UserPools[i].Id)
	}

	
	if (len(listUserPoolOutput.UserPools)) <= 1 {
		return ""
	}
	

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold | cyan }}",
		Active:   "\U0001F336 {{ . | cyan | bold }}",      // 🌶
		Inactive: "  {{ . | white }}",
		Selected: "\U0001F525 You chose: {{ . | green }}",
	}

	prompt := promptui.Select{
		Label: "Select the User Pool: ",
		Items: userPools,
		Templates: templates,
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(userPools[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
		StartInSearchMode: true,
	}

	_, userPoolId, err := prompt.Run()

	if err != nil {
		log.Print(err)
	}

	return userPoolId
}