package main

import "fmt"



func main(){
	var x string = "Hello, World"
	var b bool
	fmt.Println(x)
	boo()
	b = tonto()
	fmt.Println("The value received by tonto is:", b)
	fmt.Printf("The value received by tonto is: %t \n", b)
	cadenas("uno", "dos")
	numbers(5, 6)
	numbers(10, 14)
}

func boo(){
	fmt.Sprint("Boo function")
	fmt.Println(true && true)
}

func tonto()(bool){
	var y bool
	y = true && true
	return y
}

func cadenas(x string, y string){
	fmt.
	fmt.Printf("Hola uno es %s y dos es %s \n", x, y)
	var una, dos string =  x, y
	fmt.Println("La primera es: " + una + ", la segunda es: " + dos)
	var thisone string = fmt.Sprintf(" Con enteros %d ", 1)
	var jo string = fmt.Sprintf("jola %d", 1)
	fmt.Println("First this one" + thisone + "then: " + jo);
}

func numbers(x int, y int){
	fmt.Println("This function prints the result of an operation on screen")
	//fmt.Printf("The sum of %d + %d = %d \n", x, y, x + y)
	fmt.Printf("The sum of %d + %d = %d \n", x, y, x + y)
	fmt.Printf("The subs of %d - %d = %d \n", x, y, x - y)
}
