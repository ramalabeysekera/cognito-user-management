package helpers

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
)

func GetLocalAwsProfiles() ([]string, error) { 
	fname := config.DefaultSharedCredentialsFilename()
	if fname == ""{
		return nil, errors.New("No AWS configuration found !")
	}
	f, err := ini.Load(fname) 
	listOfProfiles := []string{}
	if err != nil {
		return nil, err
	} else {
		for _, v := range f.Sections() {
			if len(v.Keys()) != 0 {
				listOfProfiles = append(listOfProfiles, v.Name())
			}
		}
	}
	return listOfProfiles, nil
}