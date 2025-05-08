package api

import (
	"fmt"
	"log"
	"net/http"
)

func UpsBreak(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpsBreak called")
	fmt.Println("Request URL: ", r.URL.Path)
}

func UpsBreakRequest() {
	http.HandleFunc("/", UpsBreak)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	UpsBreakRequest()
}
