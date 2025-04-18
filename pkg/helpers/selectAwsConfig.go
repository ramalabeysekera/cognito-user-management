package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

// SelectAwsProfile prompts user to select an AWS profile if multiple profiles exist
// Returns the selected profile name as a string
func SelectAwsProfile() string {

	// Get list of AWS profiles from local config
	profiles, err := GetLocalAwsProfiles()

	if err != nil {
		log.Print(err)
	}

	// If multiple profiles exist, show selection prompt
	if len(profiles) > 1 {
		fmt.Println("Found multiple AWS profiles, please choose which one to use!")
		// Configure the visual templates for the selection prompt
		templates := &promptui.SelectTemplates{
			Label:    "{{ . | bold | cyan }}",
			Active:   "\U0001F336 {{ . | cyan | bold }}", // Hot pepper emoji
			Inactive: "  {{ . | white }}",
			Selected: "\U0001F525 You chose: {{ . | green }}", // Fire emoji
		}

		// Configure the selection prompt
		prompt := promptui.Select{
			Label:     "Select the AWS profile to use: ",
			Items:     profiles,
			Templates: templates,
			// Enable case-insensitive search through profiles
			Searcher: func(input string, index int) bool {
				item := strings.ToLower(profiles[index])
				input = strings.ToLower(input)
				return strings.Contains(item, input)
			},
			StartInSearchMode: true,
		}

		// Run the prompt and get selected profile
		_, profile, err := prompt.Run()

		if err != nil {
			log.Print(err)
		}

		return profile
	} else {
		// If only one profile exists, return it directly
		return profiles[0]
	}
}
