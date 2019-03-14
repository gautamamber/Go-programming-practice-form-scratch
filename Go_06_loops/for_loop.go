// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {
	// long method of loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method of loop

	fmt.Println("=================")

	for j :=1 ; j<=10; j++{
	fmt.Println(j)
}
}
