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
		okCommands := []string{"port-scanner", "port-scan", "ddos", "DDoS", "password cracker",
			"password-cracker", "password crack", "exit", "quit"}

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
	fmt.Println("Here's the result", result)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func main() {
	// var command string
	myFigure := figure.NewFigure("HACKER-ATTACKER", "poison", true)
	color.Cyan(myFigure.String())

	fmt.Println("\n\n\n\n\n\n\n")

	sessionID := getInput("Enter Session ID")
	fmt.Printf("connecting to session id: %s \n", sessionID)

	for {
		choice := prompt()
		fmt.Printf("%v", choice)
		switch choice {
		case " port-scan":
			PortScan(sessionID)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("unrecognized option")

		}

	}
	// command := getInput("Enter command")
	// fmt.Println(command)

}
