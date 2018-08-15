package main

import "fmt"



func split(sum int) (x,y,z int){
	x = sum * 4/9
	y = x * 5/8
	z = x - y
	return 
}

func main(){
	fmt.Println(split(40))
}

