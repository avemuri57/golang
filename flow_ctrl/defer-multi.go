package main

import "fmt"

func main(){
	fmt.Println("Starting Program")

	for i:=0;i<10;i++{
		defer fmt.Println(i)
		fmt.Println(i)
	}

}