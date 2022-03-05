package main

import(
	"fmt"
	"time"
	"sync"
)

func swap (first * int, second * int) {
	tmp := *first
	*first = *second
	*second = tmp
}

func bubblesort(array_slice [] int, wg *sync.WaitGroup) {
	for i:=0; i<len(array_slice); i++ {
		for j:=0; j<len(array_slice)-i-1; j++ {
			if(array_slice[j] > array_slice[j+1]) {
				swap(&array_slice[j],&array_slice[j+1])
			}
		}
	}
	wg.Done()
}

func Min(array_slice [] int) int {
	min := array_slice[0]
	for i:=0; i<len(array_slice); i++ {
		 if array_slice[i] < min {
			 min = array_slice[i]
		 }
	}

	return min
}

func main(){
	// Create Large array and fill it
	a := make([]int, 100000)
	for indx,_:= range a {
		a[indx] = 100000 - indx
	}

	// divide this large array into 4 
	a_1 := a[0:len(a)/4]
	a_2 := a[len(a)/4:len(a)/2]
	a_3 := a[len(a)/2:3*len(a)/4]
	a_4 := a[3*len(a)/4:len(a)]
	var wg sync.WaitGroup
	wg.Add(4)
	start := time.Now()
	// sort each subarray individualy
	go bubblesort(a_1,&wg)
	go bubblesort(a_2,&wg)
	go bubblesort(a_3,&wg)
	go bubblesort(a_4,&wg)
	wg.Wait()
	// merge subarrays into new array
	a_sorted := make([]int,len(a))
	a_index , a_1_index, a_2_index, a_3_index, a_4_index := 0,0,0,0,0
	for a_index != len(a) {
		elements := make([]int,0,4)
		if(a_1_index < len(a_1)) {
			elements = append(elements, a_1[a_1_index])
		}
		if(a_2_index < len(a_2)) {
			elements = append(elements, a_2[a_2_index])
		}
		if(a_3_index < len(a_3)) {
			elements = append(elements, a_3[a_3_index])
		}
		if(a_4_index < len(a_4)) {
			elements = append(elements, a_4[a_4_index])
		}

		min := Min(elements)

		if(a_1_index < len(a_1)) {
			if(a_1[a_1_index] == min){
				a_sorted[a_index] = a_1[a_1_index]
				a_1_index++
			}
		}
		if(a_2_index < len(a_2)) {
			if(a_2[a_2_index] == min){
				a_sorted[a_index] = a_2[a_2_index]
				a_2_index++
			}
		}
		if(a_3_index < len(a_3)) {
			if(a_3[a_3_index] == min){
				a_sorted[a_index] = a_3[a_3_index]
				a_3_index++
			}
		}
		if(a_4_index < len(a_4)) {
			if(a_4[a_4_index] == min){
				a_sorted[a_index] = a_4[a_4_index]
				a_4_index++
			}
		}


		a_index++
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	//fmt.Println(a_sorted)
}