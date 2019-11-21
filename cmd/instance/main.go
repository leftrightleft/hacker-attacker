package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\n")
}

// connectToSession informs the API of the connected instance
func connectToSession(string) error {
	return nil
}

func prompt() string {
	validate := func(input string) error {
		okCommands := []string{"port-scanner", "port-scan", "password cracker",
			"password-cracker", "password crack", "virus maker", "virus-maker", "virus", "ddos", "DDoS", "exit", "quit"}

		for _, v := range okCommands {
			if input == v {
				return nil
			}
		}

		return errors.New("Invalid command")
	}

	prompt := promptui.Prompt{
		Label:    "Enter Hacker-Attacker Command",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func main() {
	myFigure := figure.NewFigure("HACKER-ATTACKER", "poison", true)
	color.Red(myFigure.String())

	fmt.Print("\n\n\n\n\n\n\n")

	// TODO: Session management
	// sessionID := getInput("Enter Session ID")
	// fmt.Printf("connecting to session id: %s \n", sessionID)

	for {
		choice := prompt()

		// TODO: add a help option here
		switch choice {
		case "port-scan", "port-scanner":
			PortScan("sessionID")
		case "password-cracker", "password cracker", "password-crack":
			PasswordCracker("sessionID")
		case "virus maker", "virus-maker", "virus":
			VirusMaker("sessionID")
		case "quit", "exit":
			os.Exit(0)
		}

	}

}
