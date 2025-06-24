package main

import (
	"fmt"
	"math"
)

func main() {
	num1, num2, num3 := 45, 69, 65

	intSum := num1 + num2 + num3
	fmt.Println("Sum of the numbers: ", intSum)

	fnum1, fnum2, fnum3 := 42.3, 56.9, 85.6355

	floatSum := fnum1 + fnum2 + fnum3
	fmt.Println("Float numbers sum: ", floatSum)

	intFloatSum := num1 + int(fnum1)
	fmt.Println("intFloatsum: ", intFloatSum)

	mathSum := math.Round(floatSum*100) / 100
	fmt.Printf("The sum is now %v\n", mathSum)

	circleRadius := 15.6

	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference : %v", circumference)
}
