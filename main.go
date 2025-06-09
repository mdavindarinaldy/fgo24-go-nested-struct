package main

import "fmt"

func main() {
	type val struct {
		val string
	}

	type val1 struct {
		val val
	}

	type val3 struct {
		val []val1
	}

	type val2 struct {
		val [][]val3
	}

	str := val2{
		val: [][]val3{
			{
				{}, {}, {}, {
					val: []val1{
						{}, {
							val: val{val: "Hello"},
						},
					},
				},
			},
		},
	}

	fmt.Println(str.val[0][3].val[1].val.val)
}
