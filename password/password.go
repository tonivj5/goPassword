package password

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Characters
var (
	vocales     = "aeiouáéíóú"
	consonantes = "bcdfghjklmnñpkrstvwxyz"
	simbolos    = "!?*+-_^ "
	numeros     = "0123456789"
)

// Password is a struct that define properties of password that will be generated
type Password struct {
	vocabulary    string
	Long          int
	CanRepeatChar bool
	conf          string
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
	match, error := regexp.MatchString("@.+|[avVcCns]+", *conf)

	if error != nil {
		fmt.Println("Ha ocurrido un error, ", error)
	}

	if !match {
		fmt.Println("Opciones no válidas")

		return false
	}

	return true
}

func parseConf(conf *string) string {
	vocabulary := ""

	for _, char := range *conf {
		switch char {
		case '@':
			return string([]rune(*conf)[1:])
		case 'a':
			return (vocales + strings.ToUpper(vocales) + consonantes + strings.ToUpper(consonantes) + simbolos + numeros)
		case 'v':
			vocabulary += vocales
		case 'V':
			vocabulary += strings.ToUpper(vocales)
		case 'c':
			vocabulary += consonantes
		case 'C':
			vocabulary += strings.ToUpper(consonantes)
		case 'n':
			vocabulary += numeros
		case 's':
			vocabulary += simbolos
		}
	}

	return vocabulary
}

// IsConfigSetted return if conf has been setted
func (p *Password) IsConfigSetted() bool {
	return p.confSetted
}

// GetVocabulary return vocabulary
func (p *Password) GetVocabulary() string {
	return p.vocabulary
}
