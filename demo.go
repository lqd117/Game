package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "My Game. %s!", r.URL.Path[1:])
	if err != nil {

	}
}

func main() {
	fmt.Println(checkIdAndPassword("bingyu", "bingyu1"))
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}
