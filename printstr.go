package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, v := range s {   // v vas aller chercher tout ce qui a dans s 
		z01.PrintRune(v)   // affiche v 
	}
}
