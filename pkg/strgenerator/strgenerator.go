// Пакет strgenerator предоставляет структуру Generator, которая умеет генерировать по запросу
// строки заданной длины
package strgenerator

import (
	"math/rand"
	"strings"
)

// Предопределённые последовательности символов
const (
	LowercaseLatin = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLatin = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits         = "0123456789"
)

type Generator struct {
	chars string
	len   int

	randGen *rand.Rand
}

// NewGenerator Создаёт новый объект Generator
//
// chars - последовательность символов, из которых будут генерироваться новые строки
//
// len - длина генерируемых строк
//
// seed - сид для внутреннего генератора случайных чисел. Если он равен 0, то сид выбирается случайно
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
