// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {

	// range iterates over elements in a variety of data structures

	numbers := []int{1,2,3,4,5,6}
	sum := 0
	for _,num := range numbers{
		sum+=num
	}
	fmt.Println("Sum of numbers is : ", sum)

	for i,num := range numbers{
		if num == 3{
			fmt.Println("Index is: ", i)
		}
	}

	// Iterate range over map

	key_value := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range key_value {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range key_value {
        fmt.Println("key:", k)
    }
}