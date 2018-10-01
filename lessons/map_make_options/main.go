package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	//myGreeting := make(map[string]string){} //  shorthand
	////myGreeting := map[string]string{} // empty composite literal, shorthand
	//var myGreeting map[string]string // creates nil reference, no append
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
