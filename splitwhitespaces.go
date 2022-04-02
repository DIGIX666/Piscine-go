package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var printstr []string // on créer un tableau printstr qui contient des strings
	j := 0                // on déclare la variable j
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == ' ' { // si la string s de variable s est strictement égale à un espace
			printstr = append(printstr, s[j:i]) // dans le tableau printstr, append va ranger ce que contient s en séparant les mots
			j = i + 1                           // i + 1 car avec +2 on supprime la première lettre de tout les mots qui suis le premier et un - 1 ne peux pas être possible car sinon ça affiche la dernière lettre du mot précédent en plus du second mot
		} else if i == len(s)-1 { // sinon si i est strictement égale à la longueur total de la phrase | -1 pour avoir toute la phrase
			printstr = append(printstr, s[j:])
		}
	}
	return printstr
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
