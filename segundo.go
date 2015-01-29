package main

import "fmt"



func main(){
	var x string = "Hello, World"
	var b bool
	fmt.Println(x)
	boo()
	b = tonto()
	fmt.Println("The value received by tonto is:", b)
	cadenas("uno", "dos")
}

func boo(){
	fmt.Println(true && true)
}

func tonto()(bool){
	var y bool
	y = true && true
	return y
}

func cadenas(x string, y string){
	var una, dos string =  x, y
	fmt.Println("La primera es: " + una + ", la segunda es: " + dos)
}
