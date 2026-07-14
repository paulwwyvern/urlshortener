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
	m    map[string]string
	r    Reseter
}

func main() {
	rs := &ResetableStruct{}
	rs.Reset()
}
