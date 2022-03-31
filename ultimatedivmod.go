package piscine

func UltimateDivMod(a *int, b *int) {
	y := *a   // on déclare une variable pointeur a pour l'utiliser sur le reste

	*a = *a / *b 
	*b = y % *b // grâce à variable créée y qui récupère *a permets de rappeler *a pour faire stocké le reste dans *b
}
