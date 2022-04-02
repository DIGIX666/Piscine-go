package piscine

func ConcatParams(args []string) string {
	resultat := ""
	for i := range args { // va compter les mots stocké dans args et va incrémenter i et mots
		i++
	}

	for _, v := range args {
		resultat += "\n" + v
	}
	return resultat[1:] // on retourne le résulat du début avec 1 qui permet de partir du premier mots et le : permet de donner un résultat en string
}
