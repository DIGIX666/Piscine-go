package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // on stock le résultat de la division qui est le quotien dans le pointeur div
	*mod = a % b // on stock le reste de la division grâce au % dans le pointeur mod
}
