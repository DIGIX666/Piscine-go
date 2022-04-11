package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	x := l
	y := 0
	for x != nil && y < pos {
		y++
		x = x.Next
	}
	return x
}
