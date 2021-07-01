package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	types "github.com/hoangnguyen94/simplesurance-coding/types"
)

var t = types.Timestamps{}

func (t *types.Timestamps) Clean(current int64) {
	for len(t.data) > 0 {
		if current-t.data[0] >= 60 {
			t.data = t.data[1:]
		} else {
			break
		}
	}
	t.data = append(t.data, current)
	t.len = len(t.data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	current := time.Now().Unix()
	t.Clean(current)
	fmt.Fprintf(w, strconv.Itoa(t.len))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
