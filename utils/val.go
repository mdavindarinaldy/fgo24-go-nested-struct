package utils

type Val struct {
	Val string
}

type Val1 struct {
	Val Val
}

type Val3 struct {
	Val []Val1
}

type Val2 struct {
	Val [][]Val3
}
