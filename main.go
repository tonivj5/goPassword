package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xxxtonixxx/goPassword/generator"
	"github.com/xxxtonixxx/goPassword/password"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Este programa necesita parámetros")
		return
	}

	var (
		app         = "goPassword"
		salto       = "\t\t"
		nextLine    = "\n" + salto
		hVocabulary = "--vocabulary, -V string\t->\tVocabulario para la pass"
		hVerbose    = "--verbose, -v bool\t->\tModo verbose"
		hRepeat     = "--repeat, -r bool\t->\tRepetir caracteres"
		hLong       = "--long, -l uint\t\t->\tLongitud de la pass"
	)

	var (
		vocabulary = "a"
		long       uint
		repeat     = true
		verbose    bool
	)

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h", "--help":
			fmt.Fprintln(os.Stdout, "Uso de "+app+":"+nextLine+hVocabulary+nextLine+hVerbose+nextLine+hRepeat+nextLine+hLong)

			return
		case "--long", "-l":
			longitud, err := strconv.ParseUint(os.Args[i+1], 10, strconv.IntSize)
			if err != nil {
				fmt.Fprintln(os.Stderr, "¡Valor incorrecto en la longitud de la password!,", os.Args[i+1])

				return
			}

			long = uint(longitud)

		case "--vocabulary", "-V":
			vocabulary = os.Args[i+1]
		case "--repeat", "-r":
			if os.Args[i+1] == "false" {
				repeat = false
			}
		case "-v", "--verbose":
			verbose = true
		}
	}

	p := &password.Password{
		Long:          long,
		CanRepeatChar: repeat,
	}

	p.SetConf(vocabulary)

	pass, err := generator.GenPass(p)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error al generar la password:", err)

		return
	}

	if verbose {
		fmt.Println("Lista de parámetros: ")
		for i, arg := range os.Args {
			fmt.Fprintln(os.Stdout, i, "arg ->", arg)
		}

		fmt.Fprintln(os.Stdout, "El vocabulario será:", p.GetVocabulary())

		if repeat {
			fmt.Fprintln(os.Stdout, "Los caracteres pueden repetirse")
		} else {
			fmt.Fprintln(os.Stdout, "Los caracteres no pueden repetirse")
		}

		fmt.Println("Número de caracteres:", long)

		fmt.Fprint(os.Stdout, "La pass generada es: ")
	}

	fmt.Fprintln(os.Stdout, pass)
}
