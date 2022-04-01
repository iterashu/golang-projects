package main

import "fmt"

func main() {
	// var kanha string
	// // kanha = "30"

	// // number
	// var navindra complex64
	// // navindra = 32

	// //boolean
	// var akash bool
	// // akash = true
	// fmt.Println("kanha:==", kanha)
	// fmt.Println("navindra:==", navindra)
	// fmt.Println("akash:==", akash)

	atharva := 30
	fmt.Printf("atharva type--- %T \t value--- %v\n", atharva, atharva)
	atharva1 := 21
	atharva, mystring, mybool := 2, "ashu", true
	fmt.Printf("myint type--- %T \t value--- %v\n", atharva, atharva)
	fmt.Printf("mystring type--- %T \t value--- %v\n", mystring, mystring)
	fmt.Printf("mybool type--- %T \t value--- %v\n", mybool, mybool)
}
