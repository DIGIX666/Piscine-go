package main

import "fmt"

func ConcatParams(args []string) string {
	var resultat string // création de variale en string pour renvoyer le résultat en string
	mots := 0           // utiliser pour la taille des mots

	for i := range args { // va compter les mots stocké dans args et va incrémenter i et mots
		i++
		mots++
	}

	for i, v := range args { // v va faire une copie de tout ce que i à récupéré
		resultat += v    // va récupérer le mot
		if i != mots-2 { //
			resultat += "\n" // le résultat égale le mot plus un retour à la ligne tout de suite après
		}
	}
	return resultat // affiche le mot
}

func main() {
	test := []string{"Hello", "how", "are", "you?", "go"}
	fmt.Println(ConcatParams(test))
}
