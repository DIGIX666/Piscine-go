package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

type Door struct {
	state bool
}

// OPEN = true
// CLOSE = false

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = true
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	ptrDoor.state = false
	return false
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	ptrDoor.state = false
	return false
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = true
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if !IsDoorOpen(door) {
		CloseDoor(door)
	}
}
