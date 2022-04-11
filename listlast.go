package piscine

func ListLast(l *List) interface{} {
	x := l.Head
	for x != nil {
		if x.Next == nil {
			return x.Data
		}
		x = x.Next
	}
	return nil
}
