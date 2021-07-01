package main

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
