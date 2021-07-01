package libs

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
)

const fileName string = "timestamps.log"

func Write(t int64) {
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

func ReadFromFile() []int64 {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0644)
	var result []int64
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return result
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
