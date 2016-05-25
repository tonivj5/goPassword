package generator

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/xxxtonixxx/goPassword/password"
)

// GenPassToWriter generate a password to writer
func GenPassToWriter(w io.Writer, p *password.Password) error {
	if !p.IsConfigSetted() {
		return errors.New("Configuración sin settear")
	}

	vocabulary := []rune(p.GetVocabulary())

	if p.Long == 0 {
		p.Long = uint(len(vocabulary))
	}

	zero := 0

	rand.Seed(time.Now().Unix())
	isLimitReached := false

	for i := uint(0); i < p.Long; i++ {

		if isLimitReached {
			return errors.New("Límite alcanzado")
		}

		long := len(vocabulary)

		r := random(zero, long)
		character := vocabulary[r]
		fmt.Fprint(w, string(character))

		if !p.CanRepeatChar {
			isLimitReached = deleteRune(&vocabulary, &character)
		}
	}

	return nil
}

// GenPass generate a password
func GenPass(p *password.Password) (string, error) {
	if !p.IsConfigSetted() {
		return "", errors.New("Configuración sin settear")
	}

	vocabulary := []rune(p.GetVocabulary())

	if p.Long == 0 {
		p.Long = uint(len(vocabulary))
	}

	zero := 0
	var password bytes.Buffer

	rand.Seed(time.Now().Unix())
	isLimitReached := false

	for i := uint(0); i < p.Long; i++ {

		if isLimitReached {
			return password.String(), errors.New("Límite alcanzado")
		}

		long := len(vocabulary)

		r := random(zero, long)
		character := vocabulary[r]
		password.WriteRune(character)

		if !p.CanRepeatChar {
			isLimitReached = deleteRune(&vocabulary, &character)
		}
	}

	return password.String(), nil
}

func deleteRune(vocabulary *[]rune, character *rune) bool {
	pos := getPos(vocabulary, character)
	long := len(*vocabulary)

	if long == 1 {
		return true
	}

	(*vocabulary)[pos] = (*vocabulary)[long-1]
	(*vocabulary)[long-1] = 0
	(*vocabulary) = (*vocabulary)[:long-1]

	return false
}

func getPos(vocabulary *[]rune, character *rune) int {
	for i, char := range *vocabulary {
		if *character == char {
			return i
		}
	}

	return -1
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
