package main

//#include<stdio.h>
//void hello();

import "C"
import "fmt"

func add(x int,y int) int {
	return x + y
}

func twoReturns(x int ,y int) (int,int) {
	return 0,0;
}

func fiveReturns() (int,int,int,bool) {
	return 4,3,2,true;
}

func main(){
	C.hello();
	a := 10;
	fmt.Print(fiveReturns());
	fmt.Printf("%d\n",a);
	fmt.Print(twoReturns(123,456));

	var p int = 123;
	var s string = "\nhello world\nhello linux\n";
	fmt.Print(s);
	
	for i := 0; i < 5; i++ {
		fmt.Print(s);
	}
	fmt.Print(p);
}