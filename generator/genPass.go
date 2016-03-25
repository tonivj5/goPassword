package generator

import (
	"errors"
	"math/rand"
	"time"

	"github.com/xxxtonixxx/goPassword/password"
)

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
	password := ""

	rand.Seed(time.Now().Unix())
	isLimitReached := false

	for i := uint(0); i < p.Long; i++ {

		if isLimitReached {
			return password, errors.New("Límite alcanzado")
		}

		long := len(vocabulary)

		r := random(&zero, &long)
		character := string(vocabulary[r])
		password += character
		if !p.CanRepeatChar {
			isLimitReached = deleteRune(&vocabulary, &character)
		}
	}

	return password, nil
}

func deleteRune(vocabulary *[]rune, character *string) bool {
	pos := getPos(vocabulary, character)
	long := len(*vocabulary)
	// FIXME: delete this line fmt.Println(pos, long)

	if long == 1 {
		(*vocabulary)[0] = 0

		return true
	}

	(*vocabulary)[pos] = (*vocabulary)[long-1]
	(*vocabulary)[long-1] = 0
	(*vocabulary) = (*vocabulary)[:long-1]

	return false
}

func getPos(vocabulary *[]rune, character *string) int {
	for i, char := range *vocabulary {
		if *character == string(char) {
			return i
		}
	}

	return -1
}

func random(min, max *int) int {
	return rand.Intn(*max-*min) + *min
}
