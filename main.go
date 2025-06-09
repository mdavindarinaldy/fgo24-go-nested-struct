package main

import (
	"fgo24-nested-struct/utils"
	"fmt"
)

func main() {
	str := utils.Val2{
		Val: [][]utils.Val3{
			{
				{}, {}, {}, {
					Val: []utils.Val1{
						{}, {
							Val: utils.Val{Val: "Hello"},
						},
					},
				},
			},
		},
	}

	fmt.Println(str.Val[0][3].Val[1].Val.Val)
}
