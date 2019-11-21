package main

import (
	"fmt"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// PortScan simulates a port scanner
func PortScan(session string) {
	// TODO: Display that software is out of date
	fmt.Print("\n\n")
	figure.NewFigure("Port Scanner", "", true).Print()
	fmt.Print("\n\n")

	fmt.Print("++++++++++++++++   INITIATING SCAN   ++++++++++++++++\n\n")

	t := 0
	for ok := true; ok; ok = t != 6 {
		time.Sleep(2 * time.Second)
		fmt.Println("Scanning...  Targets found: ", t)
		t++
	}

	fmt.Print("\n\n++++++++++++++++   SCAN COMPLETE   ++++++++++++++++\n\n")
	fmt.Println(`
==================
IDENTIFIED TARGETS
==================

TARGET	        PORT	INFO
======	        ====	====
334.777.121.2	22	SSH Session
334.777.121.2	80	HTTP
334.777.121.2	443	HTTPS
334.777.121.2	8080	MYSQL
334.777.121.2	8088	Unknown
	`)
}
