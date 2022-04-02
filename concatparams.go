package piscine

func ConcatParams(args []string) string {
	resultat := ""
	for i := range args { // va compter les mots stocké dans args et va incrémenter i et mots
		i++
	}

	for _, v := range args {
		resultat += "\n" + v
	}
	return resultat[1:]
}
