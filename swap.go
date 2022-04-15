package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a // on échange les valeurs de pointeur a et pointeur b
}
