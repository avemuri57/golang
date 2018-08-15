package main

import "fmt"


// For is Golang's while loop

func main(){
	sum := 0

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

// infinite for loop with for{}