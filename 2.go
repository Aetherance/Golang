package main
 
/*
    #include <stdio.h>

    void myPrint(){
        printf("hello world!\n");
    }
*/
import "C"
 
func main() {
    C.myPrint()
}