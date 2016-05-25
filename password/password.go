package password

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Characters
var (
	vocales     = "aeiou"
	consonantes = "bcdfghjklmnpkrstvwxyz"
	simbolos    = "!?*+-_^ "
	utf8        = "áéíóúñ"
	numeros     = "0123456789"
)

// Password is a struct that define properties of password that will be generated
type Password struct {
	vocabulary    string
	Long          uint
	CanRepeatChar bool
	confSetted    bool
}

// SetConf set conf var and return true if this has been setted
func (p *Password) SetConf(options string) error {
	if validateConf(&options) {
		p.confSetted = true
		p.vocabulary = parseConf(&options)

		return nil
	}

	p.confSetted = false
	return errors.New("Configuración no válida")
}

// ValidateConf validate if configuration is okay
func validateConf(conf *string) bool {
	match, err := regexp.MatchString("@.+|[avVcCuUns]+", *conf)

	if err != nil {
		fmt.Println("Ha ocurrido un error, ", err)
	}

	if !match {
		fmt.Println("Opciones no válidas")

		return false
	}

	return true
}

func parseConf(conf *string) string {
	var vocabulary bytes.Buffer

	for _, char := range *conf {
		switch char {
		case '@':
			return (*conf)[1:]
		case 'a':
			vocabulary.WriteString(vocales)
			vocabulary.WriteString(strings.ToUpper(vocales))
			vocabulary.WriteString(consonantes)
			vocabulary.WriteString(strings.ToUpper(consonantes))
			vocabulary.WriteString(simbolos)
			vocabulary.WriteString(utf8)
			vocabulary.WriteString(numeros)
			vocabulary.WriteString(strings.ToUpper(utf8))

			return vocabulary.String()
		case 'v':
			vocabulary.WriteString(vocales)
		case 'V':
			vocabulary.WriteString(strings.ToUpper(vocales))
		case 'c':
			vocabulary.WriteString(consonantes)
		case 'C':
			vocabulary.WriteString(strings.ToUpper(consonantes))
		case 'n':
			vocabulary.WriteString(numeros)
		case 'u':
			vocabulary.WriteString(utf8)
		case 'U':
			vocabulary.WriteString(strings.ToUpper(utf8))
		case 's':
			vocabulary.WriteString(simbolos)
		}
	}

	return vocabulary.String()
}

// IsConfigSetted return if conf has been setted
func (p *Password) IsConfigSetted() bool {
	return p.confSetted
}

// GetVocabulary return vocabulary
func (p *Password) GetVocabulary() string {
	return p.vocabulary
}
