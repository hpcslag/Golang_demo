package main

import "fmt"

const (
	constStr = "123"
	constBool = true
)

func main(){
	const num int =  123
	//constBool = false <- you can't change
	fmt.Print(num);
}