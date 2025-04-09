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
	Run: func(cmd *cobra.Command, args []string) {
		userPool := selections.SelectUserPool(config.AwsConfig)
		if userPool != "" {
			permanentpassword, _ := cmd.Flags().GetBool("permanentpassword")
			attrs, err := common.DescribeUserSignInAttr(&userPool, config.AwsConfig, context.Background())

			if err != nil{
				log.Fatal(err)
			}

			if len(attrs) > 0 {
				if len(attrs) > 1 {
					selectedAttr , err := helpers.InteractiveSelection(attrs, "Please select the attribute you would like to use: ")
					if err != nil{
						log.Fatal(err)
					}

					attToFriendlyName := make(map[string](string))

					attToFriendlyName["email"] = "Email"
					attToFriendlyName["phone_number"] = "Phone Number"

					createCognitoUser(context.Background(), userPool, permanentpassword, attToFriendlyName[selectedAttr])
				}else{
					createCognitoUser(context.Background(), userPool, permanentpassword, attrs[0])
				}
			}else{
				createCognitoUser(context.Background(), userPool, permanentpassword, "")
			}
		}else{
			log.Fatal("No user pool ID found")
		}
		
	},
}

// init adds the create command to the root command
func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().Bool("permanentpassword", false, "Set password as permanant for the new user")
}


// createCognitoUser handles the creation of a new user in AWS Cognito
func createCognitoUser(ctx context.Context, userPoolId string, permpass bool, attr string){

	// Initialize Cognito client
	cogClient := cognitoidentityprovider.NewFromConfig(config.AwsConfig)
	
	// Set up input reader for user interaction
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Attempting to create the user on userPoolId:",userPoolId)
	fmt.Println("Cancel the operation if this is not intended - Ctrl+C")

	if attr != ""{
		fmt.Printf("Please enter the %v : ", attr)
	}else{
		// Get username from user input
		fmt.Print("Please enter the username: ")
	}

	userName, err := reader.ReadString('\n')
	
	if err != nil{
		log.Print(err)
	}

	userName = strings.TrimSpace(userName)

	// Get temporary password from user input
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

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Create user in Cognito
	AdminCreateUserOutput , err := cogClient.AdminCreateUser(ctx,&userInput)
	

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

	if permpass {

		_, err := common.SetPermanentPassword(userPoolId, userName, tempPassword , config.AwsConfig, ctx)

		if err != nil{
			log.Print(err)

		}else{
			log.Print("Permanant password set !")

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