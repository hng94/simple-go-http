package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const fileName string = "timestamps.log"

type timestamps struct {
	data []int64
	len  int
}

func (t *timestamps) Clean(current int64) {
	// for _, v := range t.data {
	// 	fmt.Println(v)
	// }
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

func write(t int64) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(f)
	_, _ = writer.WriteString(strconv.FormatInt(t, 10) + "\n")

	writer.Flush()
	f.Close()
}

func readFromFile() []int64 {
	file, err := os.Open(fileName)
	result := []int64{}
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, line := range txtlines {
		value, error := strconv.ParseInt(line, 10, 64)
		if error == nil {
			result = append(result, value)
		}
	}

	return result
}

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
