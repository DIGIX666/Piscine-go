package piscine

func ListPushFront(l *List, data interface{}) {
	end := &NodeL{Data: data} //  on initialise un noeud et on met nos données à l'intérieur {....}
	if l.Head == nil {        // l.head correspond au premier noeud
		l.Head = end
	} else {
		end.Next = l.Head // on dit que la variable pour passer au suivant devient le premier noeud
		l.Head = end      // le premier noeud devient le dernier
	}
}
