// Package cmd provides command-line interface functionality for managing AWS Cognito users
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
	"github.com/ramalabeysekera/cognitousermanagement/config"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/common"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/helpers"
	"github.com/ramalabeysekera/cognitousermanagement/pkg/selections"
	"github.com/spf13/cobra"
)

// createCmd represents the create command for creating new Cognito users
// It uses the cobra CLI framework to handle command execution
var createCmd = &cobra.Command{
	Use:   "createuser",
	Short: "This command is used to create a new Cognito User",
	Long: `The "createUser" command allows you to register a new user in an AWS Cognito User Pool.

Run this command with "--permanentpassword=true" to set a permanant password during the creation
Ensure your AWS credentials are properly configured before running this command.
The command uses the AWS SDK for Go (v2) and requires appropriate IAM permissions to access Cognito services`,
	// Run defines the main execution logic for the create command
	Run: func(cmd *cobra.Command, args []string) {
		// Get selected user pool from available pools
		userPool := selections.SelectUserPool(config.AwsConfig)
		if userPool != "" {
			// Check if permanent password flag is set
			permanentpassword, _ := cmd.Flags().GetBool("permanentpassword")
			// Get user sign-in attributes for the pool
			attrs, err := common.DescribeUserSignInAttr(&userPool, config.AwsConfig, context.Background())

			if err != nil{
				log.Fatal(err)
			}

			// Handle different attribute configurations
			if len(attrs) > 0 {
				if len(attrs) > 1 {
					// If multiple attributes available, let user select one
					selectedAttr , err := helpers.InteractiveSelection(attrs, "Please select the attribute you would like to use: ")
					if err != nil{
						log.Fatal(err)
					}

					// Map attribute names to friendly display names
					attToFriendlyName := make(map[string](string))
					attToFriendlyName["email"] = "Email"
					attToFriendlyName["phone_number"] = "Phone Number"

					createCognitoUser(context.Background(), userPool, permanentpassword, attToFriendlyName[selectedAttr])
				}else{
					// If only one attribute, use it directly
					createCognitoUser(context.Background(), userPool, permanentpassword, attrs[0])
				}
			}else{
				// If no attributes, create user without attribute
				createCognitoUser(context.Background(), userPool, permanentpassword, "")
			}
		}else{
			log.Fatal("No user pool ID found")
		}
		
	},
}

// init adds the create command to the root command and sets up command flags
func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().Bool("permanentpassword", false, "Set password as permanant for the new user")
}

// createCognitoUser handles the creation of a new user in AWS Cognito
// Parameters:
// - ctx: Context for AWS API calls
// - userPoolId: ID of the Cognito user pool
// - permpass: Boolean indicating if password should be permanent
// - attr: User attribute to be used (email/phone)
func createCognitoUser(ctx context.Context, userPoolId string, permpass bool, attr string){
	// Initialize Cognito client with AWS config
	cogClient := cognitoidentityprovider.NewFromConfig(config.AwsConfig)
	
	// Set up input reader for user interaction
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Attempting to create the user on userPoolId:",userPoolId)
	fmt.Println("Cancel the operation if this is not intended - Ctrl+C")

	// Prompt for username or attribute value
	if attr != ""{
		fmt.Printf("Please enter the %v : ", attr)
	}else{
		fmt.Print("Please enter the username: ")
	}

	// Read and process username input
	userName, err := reader.ReadString('\n')
	if err != nil{
		log.Print(err)
	}
	userName = strings.TrimSpace(userName)

	// Get temporary password
	fmt.Print(`Please enter the temporary password (Run this command with "--permanentpassword=true" to set a permanant password) : `)
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

	// Set timeout context for AWS API call
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Create user in Cognito
	AdminCreateUserOutput , err := cogClient.AdminCreateUser(ctx,&userInput)
	
	// Handle user creation response
	if err != nil {
		log.Fatal(err)
	} else {
		if permpass {
			log.Println("User created successfully") 
		}else{
			log.Println("User created successfully") 
			log.Printf("Username: %s, UserStatus: %s", 
			*AdminCreateUserOutput.User.Username,
			AdminCreateUserOutput.User.UserStatus)
		}
	}

	// Handle permanent password setting if requested
	if permpass {
		// Set permanent password
		_, err := common.SetPermanentPassword(userPoolId, userName, tempPassword , config.AwsConfig, ctx)

		if err != nil{
			log.Print(err)
		}else{
			log.Print("Permanant password set !")

			// Get and display updated user status
			AdminGetUserOutput, err := common.AdminGetUser(userName, userPoolId, config.AwsConfig, ctx)
			
			if err != nil{
				log.Print(err)
			}

			log.Printf("Username: %s, UserStatus: %s", 
			*AdminGetUserOutput.Username,
			AdminGetUserOutput.UserStatus)
		}
	}
}