// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
import "fmt"

// main function must, run automatically
func main() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type).	

	new_map := make(map[string]int)
	new_map["a"] = 1
	new_map["b"] = 2
	new_map["c"] = 3

	fmt.Println("Map", new_map)

	v1 := new_map["a"]
	fmt.Println("v1 : ", v1)
	fmt.Println("len is: ", len(new_map))

	// Delete

	delete(new_map, "b")
	fmt.Println(new_map)

}