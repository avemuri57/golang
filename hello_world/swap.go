package main

import "fmt"


func swap(x string, y string) (string, string){
	return y,x
}

func main(){
	fmt.Println(swap("World","Hello"))
}