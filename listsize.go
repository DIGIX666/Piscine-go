package piscine

func ListSize(l *List) int {
	x := l.Head // part du premier noeud
	count := 0
	for x != nil {
		count++
		x = x.Next // fais la même que x++
	}
	return count
}
