package strgenerator

import (
	"math/rand"
	"strings"
)

const LowercaseLatin = "abcdefghijklmnopqrstuvwxyz"
const UppercaseLatin = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "0123456789"

type Generator struct {
	chars string
	len   int

	randGen *rand.Rand
}

func NewGenerator(chars string, len int, seed int64) *Generator {
	if seed == 0 {
		seed = rand.Int63()
	}
	return &Generator{
		chars:   chars,
		len:     len,
		randGen: rand.New(rand.NewSource(seed)),
	}
}

func (g *Generator) Generate() string {
	res := strings.Builder{}

	res.Grow(g.len)

	for i := 0; i < g.len; i++ {
		res.WriteByte(g.chars[g.randGen.Intn(len(g.chars))])
	}

	return res.String()
}
