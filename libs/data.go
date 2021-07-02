package libs

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
)

func AppendFile(t int64, fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(f)
	_, _ = writer.WriteString(strconv.FormatInt(t, 10) + "\n")

	writer.Flush()
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func CleanFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
func ReadFromFile(fileName string) []int64 {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0644)
	var result []int64
	if err != nil {
		log.Fatalf("failed opening f: %s", err)
		return result
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		value, error := strconv.ParseInt(line, 10, 64)
		if error == nil {
			result = append(result, value)
		}
	}
	return result
}

func RemoveFile(fileName string) {
	e := os.Remove(fileName)
	if e != nil {
		log.Fatal(e)
	}
}
