package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL // premier noeud
	Tail *NodeL // dernier noeud
}

func ListPushBack(l *List, data interface{}) {
	first := &NodeL{Data: data} //  on initialise un noeud et on met nos données à l'intérieur {....} | tableau de donnée
	if l.Head == nil {          // l.head correspond au premier noeud
		l.Head = first
	} else {
		x := l.Head
		for x.Next != nil {
			x = x.Next // .Next veut dire le noeud suivant | du coup on  dit en partant du premier noeud tu fais la suite
		}
		x.Next = first
	}

	/*n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}*/
}
