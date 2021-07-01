package file_handler

import (
	"fmt"
	"log"
	"os"
	"bufio"
	""
)

func (t *[]int64) write() {
	f, err := os.OpenFile("timestamps.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(f)
	for _, data := range t.data
	defer f.Close()
}
