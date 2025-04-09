package helpers

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

func SelectAwsProfile() string {

	profiles, err := GetLocalAwsProfiles()

	if err != nil{
		log.Print(err)
	}

	if len(profiles) > 1 {
		fmt.Println("Found multiple AWS profiles, please choose which one to use!")
		templates := &promptui.SelectTemplates{
			Label:    "{{ . | bold | cyan }}",
			Active:   "\U0001F336 {{ . | cyan | bold }}", // 🌶
			Inactive: "  {{ . | white }}",
			Selected: "\U0001F525 You chose: {{ . | green }}",
		}
	
		prompt := promptui.Select{
			Label:     "Select the AWS profile to use: ",
			Items:     profiles,
			Templates: templates,
			Searcher: func(input string, index int) bool {
				item := strings.ToLower(profiles[index])
				input = strings.ToLower(input)
				return strings.Contains(item, input)
			},
			StartInSearchMode: true,
		}
	
		_, profile, err := prompt.Run()
	
		if err != nil {
			log.Print(err)
		}
	
		return profile
	}else{
		return profiles[0]
	}
}