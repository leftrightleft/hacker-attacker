package main

import (
	"fmt"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// PortScan simulates a port scanner
func (c colors) PortScan(session string) {
	// TODO: Display that software is out of date
	fmt.Print("\n\n")
	figure.NewFigure("Port Scanner", "", true).Print()
	fmt.Print("\n\n")

	Typer("Welcome to port-scanner.  Ports are similar to 'channels' on your computer.  The port scanner checks each of these 'channels' to see if there is an app running on that port that is out of date")
	Waiter(4)
	Typer("When port-scanner finds an out of date app, it will let you know")
	fmt.Println("\n\n")

	ip := GetInput("Enter Hackers IP Address")

	Typer("Initiating Scanner")
	Waiter(5)

	fmt.Printf("\n\n++++++++++++++++   NOW SCANNING %s  ++++++++++++++++\n\n", c.red(ip))

	t := 0
	for ok := true; ok; ok = t != 6 {
		time.Sleep(2 * time.Second)
		fmt.Println("Scanning...  Targets found: ", t)
		t++
	}

	fmt.Print("\n\n++++++++++++++++   SCAN COMPLETE   ++++++++++++++++\n\n")
	fmt.Printf(`

FOUND PORTS
==================

TARGET	        PORT	INFO
======	        ====	====
%s	23	Email      %s 
%s	80	Web Browser
%s	443	Web Browser
%s	8080	Instagram
%s	8088	SnapChat
	`, ip, c.boldRed("!!! Out of date !!!"), ip, ip, ip, ip)
}
