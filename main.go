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
		fmt.Println(help())

		os.Exit(2)
	}

	var (
		vocabulary = "a"
		long       uint
		repeat     = true
		verbose    bool
	)

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h", "--help":
			fmt.Println(help())

			return
		case "--long", "-l":
			longitud, err := strconv.ParseUint(os.Args[i+1], 10, strconv.IntSize)
			if err != nil {
				fmt.Fprintln(os.Stderr, "¡Valor incorrecto en la longitud de la password!,", os.Args[i+1])

				os.Exit(1)
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

		os.Exit(1)
	}

	if verbose {
		fmt.Println("Lista de parámetros: ")
		for i, arg := range os.Args {
			fmt.Println(i, "arg ->", arg)
		}

		fmt.Println("El vocabulario será:", p.GetVocabulary())

		fmt.Print("Los caracteres ")
		if repeat {
			fmt.Print("no ")
		}
		fmt.Println("pueden repetirse")

		fmt.Println("Número de caracteres:", long)

		fmt.Print("La pass generada es: ")
	}

	fmt.Println(pass)
}

func help() string {
	var (
		app         = "goPassword"
		salto       = "\t\t"
		nextLine    = "\n" + salto
		hVocabulary = "--vocabulary, -V string\t->\tVocabulario para la pass"
		hVerbose    = "--verbose, -v bool\t->\tModo verbose"
		hRepeat     = "--repeat, -r bool\t->\tRepetir caracteres"
		hLong       = "--long, -l uint\t\t->\tLongitud de la pass"
	)

	return "Uso de " + app + ":" + nextLine + hVocabulary + nextLine + hVerbose + nextLine + hRepeat + nextLine + hLong
}
