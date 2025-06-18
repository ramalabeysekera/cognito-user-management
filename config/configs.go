// Package config provides AWS configuration functionality
package config

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ramalabeysekera/cognito-user-management/pkg/helpers"
)

// AwsConfig holds the AWS configuration settings
var AwsConfig aws.Config

// init initializes the AWS configuration when command line arguments are provided
func init() {
	// Only load config if command line args exist
	if len(os.Args) > 1 {
		if helpers.IsBrokenGitBash() {
			fmt.Print(`
═════════════════════════════════════════════════════════
   🧊 WHOA THERE, COMMANDO!                           
                                                      
   You're using Git Bash on Windows.                  
   That's cool... until it's not.                     
                                                      
   ❌ Interactive magic? Nope                         
                                                      
   👉 RUN THIS IN POWERSHELL OR CMD/TERMINAL         
   👉 OR use: winpty ./cognitousermanagement-windows.exe      
                                                      
   Trust me, I've tried fixing this.                  
   It's not you — it's Git Bash.                      
════════════════════════════════════════════════════════
`)
			os.Exit(1)
		}
		AwsConfig = helpers.LoadAwsConfig()
	}
}
