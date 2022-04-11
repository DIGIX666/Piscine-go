package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	first := &NodeL{Data: data} //  on initialise un noeud et on met nos données à l'intérieur {....}
	if l.Head == nil {          // l.head correspond au premier noeud
		l.Head = first
	} else {
		x := l.Head
		for x.Next != nil {
			x = x.Next // .Next veut dire le noeud suivant | du coup on on dit en partant du premier noeaud tu fais la suite
		}
		x.Next = first
	}
}
