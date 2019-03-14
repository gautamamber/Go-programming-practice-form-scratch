package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {

	a,b,c := 2,3,4
	if a>b && a>c{
		fmt.Println("a is greater")
	}else if b>a && b>c{
		fmt.Println("B is greater")
	}else{
		fmt.Println("C is greater")
	}
		
}