package main

import "fmt"

//  func [name](arguments PrimiereType) returnType {
func add(x int,y int) int {
	return x + y;
}

func main(){
	fmt.Println(add(12,56));
}