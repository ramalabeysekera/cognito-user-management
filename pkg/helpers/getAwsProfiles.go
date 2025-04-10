package helpers

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
)

// GetLocalAwsProfiles retrieves a list of AWS profiles from the local credentials file
// Returns a slice of profile names and any error encountered
func GetLocalAwsProfiles() ([]string, error) {
	// Get the default AWS credentials file path
	fname := config.DefaultSharedCredentialsFilename()
	if fname == "" {
		return nil, errors.New("no aws configuration found")
	}

	// Load and parse the credentials file
	f, err := ini.Load(fname)
	// Initialize slice to store profile names
	listOfProfiles := []string{}
	if err != nil {
		return nil, err
	} else {
		// Iterate through all sections (profiles) in the credentials file
		for _, v := range f.Sections() {
			// Only include sections that have configuration keys
			if len(v.Keys()) != 0 {
				listOfProfiles = append(listOfProfiles, v.Name())
			}
		}
	}
	return listOfProfiles, nil
}
