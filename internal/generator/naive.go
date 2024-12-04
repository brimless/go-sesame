package generator

import "math/rand/v2"

type NaiveGenerator struct{}

func NewNaiveGenerator() *NaiveGenerator {
	return &NaiveGenerator{}
}

func (g *NaiveGenerator) Generate(config *GeneratorConfig) string {
	length := config.length
	useUppercase := config.useUppercase
	useNumbers := config.useNumbers
	useSymbols := config.useSymbols

	charset := Lowercase
	if useUppercase {
		charset += Uppercase
	}
	if useNumbers {
		charset += Numbers
	}
	if useSymbols {
		charset += Symbols
	}

	password := ""
	for i := 0; i < length; i++ {
		password += string(charset[rand.IntN(len(charset))])
	}

	return password
}
