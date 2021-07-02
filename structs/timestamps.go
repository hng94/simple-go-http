package structs

import (
	"fmt"
	"github.com/hoangnguyen94/simplesurance-coding/libs"
	"net/http"
	"strconv"
	"time"
)

type Timestamps struct {
	Data     []int64
	Filename string
}

func (t *Timestamps) Init(fileName string) {
	t.Filename = fileName
	t.Data = libs.ReadFromFile(t.Filename)
	current := time.Now().Unix()
	t.Clean(current)
	libs.CleanFile(t.Filename)
	for _, time := range t.Data {
		libs.AppendFile(time, t.Filename)
	}
}
func (t *Timestamps) Clean(current int64) {
	// for _, v := range t.data {
	// 	fmt.Println(v)
	// }
	index := len(t.Data)
	for i := 0; i < len(t.Data); i++ {
		if current-t.Data[i] < 60 {
			index = i
			break
		}
	}
	t.Data = t.Data[index:]
	//t.Data = append(t.Data, current)
}

func (t *Timestamps) Add(time int64) {
	t.Data = append(t.Data, time)
}

func (t *Timestamps) Handler(w http.ResponseWriter, r *http.Request) {
	current := time.Now().Unix()
	t.Clean(current)
	t.Add(current)
	fmt.Fprintf(w, strconv.Itoa(len(t.Data)))
	libs.AppendFile(current, t.Filename)
}
