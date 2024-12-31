package main

import (
	"context"
	"fmt"

	"github.com/nyaosorg/go-readline-ny"
	"github.com/sakura-aoi-ororora/go-napton"
)

const (
	astMode = iota
	evalMode
)

func main() {
	fmt.Print("Napton: Nano Lisp Implment like lepton\nPlease type 'exit' for exit repl.\n")

	var editor readline.Editor
	mode := astMode
	for {
    	code, _ := editor.ReadLine(context.Background())
		if code == "exit" {
			break
		} else if code == ":ast"{
			mode = astMode
			continue
		} else if code == ":eval" {
			mode = evalMode
			continue
		} else if code == "" {
			continue
		}

		switch mode {
		case astMode:
			node, err := napton.Parse(code)
			if err != nil {
				fmt.Printf("Error: %#v\n", err)
				continue
			} else {
				fmt.Printf("AST: %#v\n", node)
			}
		case evalMode:
			val, err := napton.NewRuntimeBuilder().Make().Run(code)
			if err != nil {
				fmt.Printf("Error: %#v\n", err)
				continue
			} else {
				fmt.Printf("Value: ")
				val.Print()
				fmt.Print("\n")
			}
		
		default:
			panic("Unknown Mode")
		}


	}

	fmt.Println("Exit.")
}