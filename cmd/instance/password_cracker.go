package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// PasswordCracker simulates a PW cracker
func (c colors) PasswordCracker(session string) {
	fmt.Print("\n\n")
	figure.NewFigure("Password Cracker", "", true).Print()
	fmt.Print("\n\n")
	Typer("Password-cracker is a tool that tries several passwords until it finds one that works.\n")
	Typer("This tools works great if the password is easy to guess.\n\n")
	_ = GetInput("Enter Hacker IP Address")
	_ = GetInput("Enter Port Number")

	fmt.Print("\n\n")
	Typer("Starting Password Cracker")
	Waiter(4)

	fmt.Print("++++++++++++++++   NOW SCANNING   ++++++++++++++++\n\n")

	t := 0
	for ok := true; ok; ok = t != 5 {
		time.Sleep(1 * time.Second)
		fmt.Println(".")
		t++

	}

	passwordList, err := readLines("./data/passwords")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	t = 0
	for ok := true; ok; ok = t != 2 {
		for _, p := range passwordList {
			time.Sleep(20 * time.Millisecond)
			fmt.Println(p)
		}
		t++
	}

	Typer("\nPassword Found")
	Waiter(5)
	fmt.Printf(`

==============================
Cracked Password:  %s
==============================

	`, c.boldRed("password123"))
}
