package helpers

import (
	"fmt"
	"log"
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
		profile := CallSingleSelect(profiles)
		return profile
	} else {
		// If only one profile exists, return it directly
		return profiles[0]
	}
}
