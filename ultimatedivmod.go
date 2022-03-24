package piscine

func UltimateDivMof(a *int, b *int) {
	*a = *a / *b
	*b = *a % *b
}
