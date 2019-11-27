package main

import (
	"errors"
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// connectToSession informs the API of the connected instance
func connectToSession(string) error {
	return nil
}

func prompt() string {
	validate := func(input string) error {
		okCommands := []string{"port-scanner", "port-scan", "password cracker",
			"password-cracker", "password crack", "virus maker", "virus-maker", "virus", "ddos", "DDoS", "exit", "quit", "help"}

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

type colors struct {
	red          func(a ...interface{}) string
	boldRed      func(a ...interface{}) string
	bgWhiteFgRed func(a ...interface{}) string
	green        func(a ...interface{}) string
}

func main() {
	myFigure := figure.NewFigure("HACKER-ATTACKER", "poison", true)
	color.Red(myFigure.String())
	fmt.Print("\n\n\n\n\n\n\n")

	c := colors{
		red:          color.New(color.FgRed).SprintFunc(),
		boldRed:      color.New(color.Bold, color.FgRed).SprintFunc(),
		bgWhiteFgRed: color.New(color.Bold, color.FgRed, color.BgWhite).SprintFunc(),
		green:        color.New(color.FgGreen).SprintFunc(),
	}

	Typer("Welcome to Hacker-Attacker")
	Waiter(4)
	fmt.Print("\n")

	Typer("These are the commands you can use with Hacker-Attacker.  Remember, you can view this menu at any time by typing 'help'")
	fmt.Print("\n\n")

	help()
	fmt.Print("\n\n")

	Typer("Let's get started")
	Waiter(4)
	// TODO: Session management
	// sessionID := getInput("Enter Session ID")
	// fmt.Printf("connecting to session id: %s \n", sessionID)

	for {
		choice := prompt()

		// TODO: add a help option here
		switch choice {
		case "port-scan", "port-scanner":
			tm.Clear()
			tm.MoveCursor(0, 0)
			tm.Flush()
			c.PortScan("sessionID")
			// myFigure := figure.NewFigure("HACKER-ATTACKER", "poison", true)
			// color.Red(myFigure.String())
			// fmt.Print("\n\n\n\n\n\n\n")
		case "password-cracker", "password cracker", "password-crack":
			tm.Clear()
			tm.MoveCursor(0, 0)
			tm.Flush()
			c.PasswordCracker("sessionID")
		case "virus maker", "virus-maker", "virus":
			tm.Clear()
			tm.MoveCursor(0, 0)
			tm.Flush()
			c.VirusMaker("sessionID")
		case "help":
			help()
		case "quit", "exit":
			tm.Clear()
			tm.MoveCursor(0, 0)
			tm.Flush()
			os.Exit(0)
		}

	}

}
