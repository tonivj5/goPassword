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
		directout  = false
	)

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--help", "-h":
			fmt.Println(help())

			return
		case "--long", "-l":
			if len(os.Args) < 3 {
				fmt.Println(help())

				os.Exit(1)
			}

			longitud, err := strconv.ParseUint(os.Args[i+1], 10, strconv.IntSize)
			if err != nil {
				fmt.Fprintln(os.Stderr, "¡Valor incorrecto en la longitud de la password!,", os.Args[i+1])

				os.Exit(1)
			}

			long = uint(longitud)

		case "--vocabulary", "-V":
			if len(os.Args) < 3 {
				fmt.Println(help())

				os.Exit(1)
			}
			vocabulary = os.Args[i+1]
		case "--norepeat", "-n":
			repeat = false
		case "--verbose", "-v":
			verbose = true
		case "-d", "--directout":
			directout = true
		case "-" + os.Args[i][1:]:
			fmt.Fprintln(os.Stderr, "¡Argumento no válido!", os.Args[i])

			os.Exit(1)
		}
	}

	p := &password.Password{
		Long:          long,
		CanRepeatChar: repeat,
	}

	p.SetConf(vocabulary)

	var pass string
	var err error

	if !directout {
		pass, err = generator.GenPass(p)
	}

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
		if !repeat {
			fmt.Print("no ")
		}
		fmt.Println("pueden repetirse")

		fmt.Println("Número de caracteres:", long)

		fmt.Print("La pass generada es: ")
	}

	if !directout {
		fmt.Println(pass)
	} else {
		err = generator.GenPassToWriter(os.Stdout, p)

		if err != nil {
			fmt.Fprintln(os.Stderr, "\nError al generar la password:", err)

			os.Exit(1)
		}
	}
}

func help() string {
	var (
		app            = "goPassword"
		salto          = "\t\t"
		nextLine       = "\n" + salto
		hVocabulary    = "--vocabulary, -V string\t->\tVocabulario para la pass. Las opciones disponibles son:"
		hVocabularyOpt = "\t\t\t\t\t@.+|[avVcCuUns]+ (regex)." + nextLine + "\t\t\t\t\tPor ejemplo: nVv -> 0-9 AEIOU aeiou"
		hVerbose       = "--verbose, -v\t\t->\tModo verbose"
		hRepeat        = "--norepeat, -n\t\t->\tSin repetición de caracteres"
		hLong          = "--long, -l uint\t\t->\tLongitud de la pass"
		hDirectout     = "--directout, -d\t\t->\tLa password se muestra por salida estándar conforme se va generando" + nextLine + "\t\t\t\t(si ocurre algún error como por ejemplo, en la longitud, habrá parte de la pass que se haya generado)"
	)

	return "Uso de " + app + ":" + nextLine + hVocabulary + nextLine + hVocabularyOpt + nextLine +
		hVerbose + nextLine + hRepeat + nextLine + hLong + nextLine + hDirectout
}
