package helpers

import (
	"github.com/fatih/color"
	"log"
)

func PrintSuccessLog(message string) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Print(green(message))
}

func PrintFatalErrorLog(message string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Fatal(red(message))
}

func PrintWarningErrorLog(message string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Print(red(message))
}