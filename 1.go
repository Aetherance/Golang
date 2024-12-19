package main

import "fmt"

func add(x int,y int) int {
	return x + y
}

func main() {
	var x = 1
	var y = 2
	var a , b , c = 1 , 2 , 3
	fmt.Println(a,"\n",b,"\n",c)
	fmt.Println(add(x,y))
	fmt.Println("Hello World\n"+"Hello Linux",x)
}