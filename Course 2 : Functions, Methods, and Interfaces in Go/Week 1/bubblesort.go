package main

import (
	"fmt"
)

func swap (first * int, second * int) {
	tmp := *first
	*first = *second
	*second = tmp
}

func bubblesort(array_slice [] int) {
	for i:=0; i<len(array_slice); i++ {
		for j:=0; j<len(array_slice)-i-1; j++ {
			if(array_slice[j] > array_slice[j+1]) {
				swap(&array_slice[j],&array_slice[j+1])
			}
		}
	}
}
func main() {
	a := make([]int, 0)
	var tmp int
	const max_size = 10
	for indx:=0; indx<max_size; indx++ {
		if _, err := fmt.Scanln(&tmp); err == nil {
			a = append(a,tmp)
		} else {
			break
		}
	}
	bubblesort(a[0:len(a)])
	fmt.Println(a)
}