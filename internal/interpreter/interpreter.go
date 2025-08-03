package interpreter

import (
	"fmt"

	"github.com/Jablinelang/Jabline/internal/ast"
	"github.com/Jablinelang/Jabline/internal/object"
)

var Debug bool = false

type Environment struct {
	store map[string]object.Object
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]object.Object)}
}

func Eval(node ast.Node, env *Environment) object.Object {
	if Debug {
	fmt.Printf("Evaluando nodo: %T\n", node)
}
	switch node := node.(type) {

	case *ast.Program:
		var result object.Object
		for _, stmt := range node.Statements {
			result = Eval(stmt, env)
		}
		return result

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		env.store[node.Name.Value] = val
		return val

	case *ast.EchoStatement:
		val := Eval(node.Value, env)
		fmt.Println(val.Inspect()) // ← aquí se imprime
		return val


	case *ast.Identifier:
		if val, ok := env.store[node.Value]; ok {
			return val
		}
		return object.NewError("variable no definida: " + node.Value)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.IntegerLiteral:
		return &object.Integer{Value: int64(node.Value)}

	default:
		return nil
	}
}
