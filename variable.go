//Use var keyword and = assignment operator

package main

import "fmt"

func main() {
	// Explicitly Typed Declarations

	/* var number int
	   number = 9
	     or

	var number int = 9
	fmt.Println(number)

	//Constants

	const number int = 42
	*/
	//Implicitly Typed Declaration

	number := 9
	myString := "This is string"
	fmt.Println(number)
	fmt.Println(myString)

	//Constants

	const aString = "This is another string"
	fmt.Println(aString)

}
