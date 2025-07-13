package main

import "fmt"

type box struct {
	height int
	width  int
	length int
}

func (b *box) canStack(bottom *box) bool {
	if b.height < bottom.height && b.width < bottom.width && b.length < bottom.length {
		return true
	}
	return false
}

var d []int

func boxes(b []*box) int {
	n := len(b)
	d = make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = b[i].height
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if b[j].canStack(b[i]) {
				d[i] = max(d[i], d[j]+b[i].height)
				res = max(res, d[i])
			}
		}
	}
	return res
}

func main() {
	b := []*box{
		{
			length: 2,
			width:  1,
			height: 2,
		},
		{
			length: 3,
			width:  2,
			height: 3,
		},
		{
			length: 2,
			width:  2,
			height: 8,
		},
		{
			length: 2,
			width:  3,
			height: 4,
		},
		{
			length: 2,
			width:  2,
			height: 1,
		},
		{
			length: 4,
			width:  4,
			height: 5,
		},
	}
	fmt.Println(boxes(b))
}
