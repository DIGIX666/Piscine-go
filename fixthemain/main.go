package piscine

import "github.com/01-edu/z01"

type door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Clossing ...")
	state = CLOSE
	return true
}

func IsDoorOpen(Door Door) {
	PrintStr("is the Door opened ?")
	return Door.state = OPEN
}

func IsDoorClose(ptrDoor *Door) {
	PrintStr("is the Door closed ?")
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
