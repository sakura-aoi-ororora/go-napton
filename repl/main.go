package main

import (
	"context"
	"fmt"

	"github.com/nyaosorg/go-readline-ny"
	"github.com/sakura-aoi-ororora/go-napton"
)

func main() {
	fmt.Print("Napton: Nano Lisp Implment like lepton\nPlease type 'exit' for exit repl.\n")

	var editor readline.Editor
	for {
    	code, _ := editor.ReadLine(context.Background())
		if code == "exit" {
			break
		} else if code == "" {
			continue
		}

		node, err := napton.Parse(code)
		if err != nil {
			fmt.Printf("Error: %#v\n", err)
		} else {
			fmt.Printf("AST: %#v\n", node)
		}
	}

	fmt.Println("Exit.")
}