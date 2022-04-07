package main

import "fmt"

func Compact(ptr *[]string) int {
	size := 0
	for _, v := range *ptr {
		if v != "" {
			size++
		}
	}
	count := 0
	array := make([]string, size)
	for _, r := range *ptr {
		if r != "" {
			array[count] = r
			count++
		}
	}
	return size
}

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
