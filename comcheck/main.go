package main

import (
	"fmt"
	"os"
)

func main() {
	A := os.Args[1:]
	for _, val := range A {
		if val == "01" || val == "galaxy" || val == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
