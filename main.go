package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var t types.timestamps

func handler(w http.ResponseWriter, r *http.Request) {
	current := time.Now().Unix()
	t.Clean(current)
	fmt.Fprintf(w, strconv.Itoa(t.len))
}

func main() {
	t = types.timestamps{
		data: []int64{},
		len:  0,
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
