package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Menetapkan handler untuk route "/" yang akan menampilkan "Hello, World!"
	http.HandleFunc("/", helloWorldHandler)

	// Memulai web server pada port 8080
	fmt.Println("Server berjalan pada http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
