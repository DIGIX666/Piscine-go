package piscine

func UltimateDivMod(a *int, b *int) {
	var a := *a
	var b := *b

	*a = a / b
	*b = a % b
}
