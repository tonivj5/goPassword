package main

import (
	"fmt"

	"github.com/xxxtonixxx/goPassword/generator"
	"github.com/xxxtonixxx/goPassword/password"
)

func main() {
	p := &password.Password{
		Long:          10,
		CanRepeatChar: true,
	}

	configs := []string{"Password con diccionario de caracteres personalizado", "@Toni Villena", "Password usando configuraciones", "cnV"}

	for i := 1; i < len(configs); i += 2 {
		p.SetConf(configs[i])

		pass, error := generator.GenPass(p)

		if error != nil {
			fmt.Println("Ocurrió un error:", error)
		}

		fmt.Println(configs[i-1]+":", pass)
	}
}
