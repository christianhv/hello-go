package main

import "fmt"

 var globalVar = "This is a global variable and can be used in any function"

func main(){
	var x string = "Hello, World"
	var b bool
	fmt.Println(x)
	boo()
	b = tonto()
	fmt.Println("The value received by tonto is:", b)
	fmt.Printf("The value received by tonto is: %t \n", b)
	cadenas("uno", "dos")
	numbers(10, 5)
	numbers(20, 14)
	fmt.Println("\n", globalVar)
}

func boo(){
	fmt.Println("#### Boo function ####")
	fmt.Println(true && true)
	fmt.Println("\n", globalVar)
}

func tonto()(bool){
	fmt.Println("\n#### Tonto function ####")
	var y bool
	y = true && true
	fmt.Println("\n", globalVar)
	return y
}

func cadenas(x string, y string){
	fmt.Println("\n#### Cadenas function ####")
	fmt.Printf("Hola uno es %s y dos es %s \n", x, y)
	var una, dos string =  x, y
	fmt.Println("La primera es: " + una + ", la segunda es: " + dos)
	var thisone string = fmt.Sprintf(" Con enteros %d ", 1)
	var jo string = fmt.Sprintf("jola %d", 1) 
	fmt.Println("First this one" + thisone + "then: " + jo);
	fmt.Printf("Is string \"%s\" equal to \"%s\" = %t \n", x, y, x == y)
	fmt.Println("\n", globalVar)
}

func numbers(x int, y int){
	fmt.Println("\n#### numbers function ####")
	fmt.Println("This function prints the result of an operation on screen")
	//fmt.Printf("The sum of %d + %d = %d \n", x, y, x + y)
	fmt.Printf("The sum  of %d + %d = %d \n", x, y, x + y)
	fmt.Printf("The subs of %d - %d = %d \n", x, y, x - y)
	fmt.Printf("The mult of %d * %d = %d \n", x, y, x * y)
	fmt.Printf("The div  of %d / %d = %d \n", x, y, x / y)
	fmt.Printf("The rem  of %d %% %d = %d \n", x, y, x % y)
	
	mew := 5
	fmt.Println(mew)
	fmt.Println("\n", globalVar)
}

func askAndGive{
	
}

