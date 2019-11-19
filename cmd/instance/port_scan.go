package main

import (
	"fmt"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func PortScan(session string) {
	figure.NewFigure("\nPort Scanner\n", "", true).Print()
	t := 0
	for ok := true; ok; ok = t != 6 {
		time.Sleep(2 * time.Second)
		fmt.Println("scanning...  Targets found: ", t)
		t++
	}

	fmt.Println(`
TARGET	        PORT	INFO
======	        ====	====
334.777.121.2	22		SSH Session
334.777.121.2	80		HTTP
334.777.121.2	443		HTTPS
334.777.121.2	8080	MYSQL
334.777.121.2	8088	Unknown
	`)
}
