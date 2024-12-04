package generator

import "errors"

type GeneratorConfig struct {
	length       int
	useUppercase bool
	useNumbers   bool
	useSymbols   bool
}

func NewGeneratorConfig(length int, useUppercase bool, useNumbers bool, useSymbols bool) (*GeneratorConfig, error) {
	if length < 6 || length > 128 {
		return nil, errors.New("Length must be between 6 and 128")
	}

	return &GeneratorConfig{
		length:       length,
		useUppercase: useUppercase,
		useNumbers:   useNumbers,
		useSymbols:   useSymbols,
	}, nil
}

type Generator interface {
	Generate(config *GeneratorConfig) string
}

type GeneratorType int

const (
	Naive GeneratorType = iota
)

const (
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbols   = "!@#$%^&*()-=_+[]{}|;:',.<>?/`~"
	Numbers   = "0123456789"
)

func GeneratePassword(config *GeneratorConfig, generatorType GeneratorType) string {
	var generator Generator
	switch generatorType {
	case Naive:
		generator = NewNaiveGenerator()
	// default to naive way to generate password
	default:
		generator = NewNaiveGenerator()
	}
	return generator.Generate(config)
}
