package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	tm "github.com/buger/goterm"
	"github.com/common-nighthawk/go-figure"
)

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\n")
}

func help() {
	fmt.Print("\n")
	totals := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(totals, "Command\tWhat It Does\n")
	fmt.Fprintf(totals, "-------\t------------\n")
	fmt.Fprintf(totals, "%s\t%s\n", "port-scan", "Scan the hackers computer to find a vulnerable app")
	fmt.Fprintf(totals, "%s\t%s\n", "virus-maker", "Create a virus then install it on the hackers computer")
	fmt.Fprintf(totals, "%s\t%s\n", "password-cracker", "Find the hackers password")
	fmt.Fprintf(totals, "%s\t%s\n", "ddos", "Execute a DDoS attack on a hacker")
	tm.Println(totals)

	tm.Flush()
}

// Typer - Give me your strings and I'll type them bad boys out
func Typer(str string) {
	s := strings.Split(str, "")
	for _, x := range s {
		time.Sleep(40 * time.Millisecond)
		fmt.Print(x)
	}
}

func Waiter(duration int) {
	n := 0
	for n < duration {
		n++
		time.Sleep(750 * time.Millisecond)
		fmt.Print(".")
		if n == duration {
			fmt.Print("\n")
		}
	}
}

func PrintTitle(title string) {
	fmt.Print("\n\n")
	figure.NewFigure(title, "", true).Print()
	fmt.Print("\n\n")
}

// func welcome() {
// 	welcomeMessage := "Welcome to Hacker-Attacker."
// 	s := strings.Split(welcomeMessage, "")
// 	for _, x := range s {
// 		time.Sleep(40 * time.Millisecond)
// 		fmt.Print(x)
// 	}

// 	elipsis := "...."
// 	e := strings.Split(elipsis, "")
// 	for _, y := range e {
// 		fmt.Print(y)
// 	}
// }
