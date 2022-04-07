package main

import "fmt"

type pilot struct {
	Name     string
	Age      int
	Life     int
	Aircraft int
}

func main() {
	donnie := pilot{}

	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = 1

	fmt.Println(donnie)
}
