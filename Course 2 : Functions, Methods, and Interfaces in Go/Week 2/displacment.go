package main

import (
	"fmt"
)

func GenDisplaceFn(a,v0,s0 float64) func(float64) float64{
	return func(t float64) float64 {return 0.5*a*t*t+v0*t+s0}
}

func main () {
	var a,v0,s0,t float64
	fmt.Println("Please Enter Acceleration")
	fmt.Scan(&a)
	fmt.Println("Please Enter initial velocity")
	fmt.Scan(&v0)
	fmt.Println("Please Enter initial displacement")
	fmt.Scan(&s0)
	fmt.Println("Please Enter time")
	fmt.Scan(&t)

	DisplaceFn:= GenDisplaceFn(a,v0,s0)
	fmt.Println(DisplaceFn(t))
}