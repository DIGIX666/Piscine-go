/*package main

import "fmt"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}*/

package main

import "github.com/01-edu/z01"

/*------------------------------------------\
  Comme explique nous allons donner un type
  à notre point.
\------------------------------------------*/

type point struct {
	x string
	y string
}

/*------------------------------------------\
  Nous donner une valeur à nos pointeurs
\------------------------------------------*/

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

/*-------------------------------------------------------------------------------------------------------------\
   on transforme notre s string en rune et
   après on dit que pour i égale à 0; i inférieur à len(sRune) donc la longueur de (sRune) alors i ingrémente
\-------------------------------------------------------------------------------------------------------------*/

func printStr(s string) {
	sRune := []rune(s)
	for i := 0; i < len(sRune); i++ {
		z01.PrintRune(sRune[i])
	}
}

/*--------------------------------------------------------------------------------------\
  On va déclarer une variale qui réprend ce qui a été stocker dans pointeur de point.
  On appelle la fonction setPoint avec la copie du nouveau pointeur (&points)
  En suite on déclare une variale ou on indique la forme qu'on veut pour note résultat
\--------------------------------------------------------------------------------------*/

func main() {
	points := point{}
	setPoint(&points)
	result := "x = " + points.x + ", y = " + points.y
	printStr(result)
	z01.PrintRune('\n')
}
