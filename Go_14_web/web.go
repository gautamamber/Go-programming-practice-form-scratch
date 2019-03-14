// main package
package main 

// Package fmt implements formatted I/OPackage fmt implements formatted I/O
// import web library
import ("fmt"
		"net/http")

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Hello world</h1>")
}
func new_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>this is new page</h1>")
}

// main function must, run automatically
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/new_page", new_page)
	fmt.Println("Server is listening")
	http.ListenAndServe(":3000", nil)
}