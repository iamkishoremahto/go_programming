package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	now := time.Now()
	fmt.Print("Time now: ", now, "\n")

	//Format time

	fmt.Println(now.Format(time.ANSIC) + "\n")
}
