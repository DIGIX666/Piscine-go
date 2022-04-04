package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	lenght := 0
	// ici on va stocker les arguments après le go run, pour compter le nombre d'arguments
	for i := range args {
		lenght = i + 1
	}

	if lenght > 1 { // si la longueur est supérieur à deux arguments alors c'est trop long
		fmt.Println("Too many arguments")
	} else if lenght == 0 { // si y a pas d'argument donc pas non de fichier il se passe rien
		fmt.Println("File name missing")
	} else if args[0] == "quest8.txt" { // si args[0] (correspond à index 0) est strictement égale à "quest8.txt"

		content, err := ioutil.ReadFile(args[0]) // permet de lire le fichier
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(content))
	}
}
