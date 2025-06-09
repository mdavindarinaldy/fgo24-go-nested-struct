package main

import "fmt"

func main() {
	type val struct {
		val []val
	}

	type val2 struct {
		val [][]val
	}

	str := val2{
		val: [][]val{
			{
				{}, {}, {
					val: []val{
						{}, {
							val: []val{
								// {val: "Hello!"},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(str)
}
