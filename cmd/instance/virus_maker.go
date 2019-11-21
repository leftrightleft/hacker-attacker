package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt, "-> ")
	text, _ := reader.ReadString('\n')
	return text
}

func VirusWatcher() {
	websites := []string{"howtohack.com", "cnn.com", "news.com", "www.bloomberg.com/ffj/article/?name=thisisit", "smellyfarts.com", "news.com", "mylittlepony.com", "hackerfacebook.com"}
	time.Sleep(2 * time.Second)
	fmt.Print("Virus successfully installed on hacker computer. \nNow monitoring websites the hacker is visiting...")
	time.Sleep(5 * time.Second)
	fmt.Println(
		`

===============================
WEBSITES THE HACKER IS VISITING
===============================

	`)
	for _, w := range websites {
		time.Sleep(2 * time.Second)
		fmt.Println(w)
	}
	return

}

// VirusMaker simulates generating and installing a virus
func VirusMaker(session string) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Print("\n\n")
	figure.NewFigure("Virus Maker", "", true).Print()
	fmt.Print("\n\n")
	// TODO: add a description of the virus.  It returns the Websites that the hacker is visiting
	ip := readInput("Enter Hacker IP Address")
	_ = readInput("Enter Hacker Password")

	fmt.Printf("Now generating virus for hacker IP address: %s ", red(ip))
	//  TODO: Here we need to make a progress bar of some sort
	time.Sleep(3 * time.Second)
	fmt.Println(green("SUCCESS!"))

	answer := readInput("Would you like to install a virus on the hackers computer? [Y/N]")
	if strings.ToLower(answer) == "y\n" {
		fmt.Printf("Now installing virus for hacker IP address: %s ", red(ip))
		fmt.Println("Please Wait")
		t := 0
		for ok := true; ok; ok = t != 5 {
			time.Sleep(1 * time.Second)
			fmt.Println(".")
			t++
		}

		fmt.Println(green("Success!"))
	} else {
		return
	}

	VirusWatcher()

	return

}
