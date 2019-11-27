package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// func readInput(prompt string) string {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print(prompt, "-> ")
// 	text, _ := reader.ReadString('\n')
// 	return text
// }

func VirusWatcher() {
	websites := []string{"howtohack.com", "cnn.com", "news.com", "www.bloomberg.com/ffj/article/?name=thisisit", "https://leftrightleft.github.io/hacker-site/login/", "news.com", "mylittlepony.com"}
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
func (c colors) VirusMaker(session string) {

	fmt.Print("\n\n")
	figure.NewFigure("Virus Maker", "", true).Print()
	fmt.Print("\n\n")
	Typer("Virus Maker generates and installs a virus on the hackers computer.\n")
	Typer("Once the virus is installed, Virus Maker will print which websites the hacker is visiting.\n\n")
	ip := GetInput("Enter Hacker IP Address")
	_ = GetInput("Enter Port Number")
	_ = GetInput("Enter Hacker Password")

	// // TODO: add a description of the virus.  It returns the Websites that the hacker is visiting
	// ip := readInput("Enter Hacker IP Address")
	// _ = readInput("Enter Hacker Password")

	fmt.Printf("\nNow generating virus for hacker IP address: %s\n", c.red(ip))
	//  TODO: Here we need to make a progress bar of some sort
	Typer("Please wait")
	Waiter(6)

	answer := GetInput("Would you like to install a virus on the hackers computer? [Y/N]")
	if strings.ToLower(answer) == "y" {
		fmt.Printf("Now installing virus for hacker IP address: %s ", c.red(ip))
		fmt.Println("Please Wait")
		t := 0
		for ok := true; ok; ok = t != 5 {
			time.Sleep(1 * time.Second)
			fmt.Println(".")
			t++
		}

		fmt.Println(c.green("Success!"))
	} else {
		return
	}

	VirusWatcher()

	return

}
