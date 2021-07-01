package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const fileName string = "timestamps.log"

var t = timestamps{}

func handler(w http.ResponseWriter, r *http.Request) {
	current := time.Now().Unix()
	t.Clean(current)
	fmt.Fprintf(w, strconv.Itoa(t.len))
	write(current)
}

func main() {
	t.data = readFromFile()
	t.len = len(t.data)
	current := time.Now().Unix()
	t.Clean(current)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
