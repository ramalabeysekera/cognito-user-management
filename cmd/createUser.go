package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/selections"
	"github.com/ramalabeysekera/cognitousermanagement/config"
	"github.com/spf13/cobra"
)

// createCmd represents the create command for creating new Cognito users
// It uses the cobra CLI framework to handle command execution
var createCmd = &cobra.Command{
	Use:   "createuser",
	Short: "This command is used to create a new Cognito User",
	Long: `The "createUser" command allows you to register a new user in an AWS Cognito User Pool.

Ensure your AWS credentials are properly configured before running this command.
The command uses the AWS SDK for Go (v2) and requires appropriate IAM permissions to access Cognito services`,
	Run: func(cmd *cobra.Command, args []string) {
		userPool := selections.SelectUserPool(config.AwsConfig)
		if userPool != "" {
			createCognitoUser(context.Background(), userPool)
		}else{
			log.Fatal("No user pool ID found")
		}
		
	},
}

// init adds the create command to the root command
func init() {
	rootCmd.AddCommand(createCmd)
}


// createCognitoUser handles the creation of a new user in AWS Cognito
func createCognitoUser(ctx context.Context, userPoolId string){

	// Initialize Cognito client
	cogClient := cognitoidentityprovider.NewFromConfig(config.AwsConfig)
	
	// Set up input reader for user interaction
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Attempting to create the user on userPoolId:",userPoolId)
	fmt.Println("Cancel the operation if this is not intended - Ctrl+C")

	// Get username from user input
	fmt.Print("Please enter the username: ")
	userName, err := reader.ReadString('\n')
	
	if err != nil{
		log.Print(err)
	}

	userName = strings.TrimSpace(userName)

	// Get temporary password from user input
	fmt.Print(`Please enter the temporary password (Run this app with "setpassword" command to set a permanant password) : `)
	tempPassword, err := reader.ReadString('\n')

	tempPassword = strings.TrimSpace(tempPassword)

	if err != nil{
		log.Print(err)
	}

	// Prepare user creation input parameters
	userInput := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId: &userPoolId,
		Username: &userName,
		MessageAction: types.MessageActionTypeSuppress,
		TemporaryPassword: &tempPassword,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Create user in Cognito
	AdminCreateUserOutput , err := cogClient.AdminCreateUser(ctx,&userInput)
	

	if err != nil {
		log.Print(err)
	} else {
		log.Println("User created successfully") 
		log.Printf("Username: %s, UserStatus: %s", 
		*AdminCreateUserOutput.User.Username,
		AdminCreateUserOutput.User.UserStatus)
	}}