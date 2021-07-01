package main

import (
	"fmt"
	"github.com/hoangnguyen94/simplesurance-coding/libs"
	"log"
	"net/http"
	"strconv"
	"time"
)

var t = libs.Timestamps{}

func handler(w http.ResponseWriter, r *http.Request) {
	current := time.Now().Unix()
	t.Clean(current)
	t.Data = append(t.Data, current)
	t.Len = len(t.Data)
	fmt.Fprintf(w, strconv.Itoa(t.Len))
	libs.Write(current)
}

func main() {
	t.Data = libs.ReadFromFile()
	t.Len = len(t.Data)
	current := time.Now().Unix()
	t.Clean(current)
	fmt.Println(t.Data)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
