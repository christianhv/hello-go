package main

import "fmt"



func main(){
	var x string = "Hello, World"
	var b bool = false
	fmt.Println(x)
	boo()
	b = tonto()
	fmt.Println(b)
}

func boo(){
	fmt.Println(true && true)
}

func tonto(){
	var y bool
	y = true && true
	
	return y
}
