package libs

type Timestamps struct {
	Data []int64
	Len  int
}

func (t *Timestamps) Clean(current int64) {
	// for _, v := range t.data {
	// 	fmt.Println(v)
	// }
	for len(t.Data) > 0 {
		if current-t.Data[0] >= 60 {
			t.Data = t.Data[1:]
		} else {
			break
		}
	}
	//t.Data = append(t.Data, current)
	t.Len = len(t.Data)
}
