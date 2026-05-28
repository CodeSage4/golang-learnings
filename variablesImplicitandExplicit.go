package main

import "fmt"

func main(){
	// var x int8 
	// x=2
	// this is explicit assignment of var
	y := 3 
	// this is implicit assignment of var and := is called assignment var
	// the advantage here is we can decide the type later, useful when using complex datatstructures
	// fmt.Print(x," ",y)
	fmt.Printf("%T",y)
	// this gives the type of var

	var z = uint(0)
	// here we are doing typecasting from in to uint
	fmt.Printf("%T",z)


}	