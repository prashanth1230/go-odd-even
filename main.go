package main

import "fmt"

type flt func(int) bool // aliasing type

func isEven(n int) bool { // check if n is even or not
	return n%2 == 0
}

func filter(sl []int, f flt) (yes, no []int) { // split s into two slices: even and odd
	for _, e := range sl {
		if f(e) {
			yes = append(yes, e)
		} else {
			no = append(no, e)
		}
	}
	return
}

func main() {
	val1 := []int{1, 2, 3, 4}
	even, odd := filter(val1, isEven)
	fmt.Println(even)
	fmt.Println(odd)
}
