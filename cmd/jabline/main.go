package main

import (
	"fmt"
	"os"

	"github.com/Jablinelang/Jabline/internal/interpreter"
	"github.com/Jablinelang/Jabline/internal/lexer"
	"github.com/Jablinelang/Jabline/internal/parser"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Uso: jabline run archivo.jab")
		os.Exit(1)
	}

	filePath := os.Args[2]

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error leyendo el archivo: %s\n", err)
		os.Exit(1)
	}

	// Crear lexer y parser
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	// Verificar errores de parsing
	if len(p.Errors()) > 0 {
		fmt.Println("Errores de parsing:")
		for _, e := range p.Errors() {
			fmt.Println("\t" + e)
		}
		os.Exit(1)
	}

	// Crear entorno de ejecución y evaluar
	env := interpreter.NewEnvironment()
	interpreter.Eval(program, env) // ← No imprimimos aquí, echo lo hace internamente
}
