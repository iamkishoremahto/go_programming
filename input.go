package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("This is your entered string: ", str)

	fmt.Print("Enter a number: ")
	str2, _ := reader.ReadString('\n')
	floatNumber, err := strconv.ParseFloat(strings.TrimSpace(str2), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Value of number: ", floatNumber)
	}

}
