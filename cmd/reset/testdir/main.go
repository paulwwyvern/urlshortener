package main

type Reseter interface {
	R()
}

// generate:reset
type ResetableStruct struct {
	i    int
	str  string
	strP *string
	s    []int
	a    [10]int
	m    *map[string]string
	r    ***ResetableStruct
}

func main() {
	rs := &ResetableStruct{}
	rs.Reset()
}
