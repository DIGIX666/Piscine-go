package piscine

func ListReverse(l *List) {
	x := l.Head
	var next *NodeL
	var first *NodeL

	for x != nil {
		next, x.Next = x.Next, first
		first, x = x, next
	}
	l.Head = first
}
