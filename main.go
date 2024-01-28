package main

import ("fmt" 
		"net/http")

func home(page http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(page, "Hello world !")
}

func main() {
	fmt.Println("Hello World !")

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}