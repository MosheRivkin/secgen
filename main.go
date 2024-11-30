package main

import (
	"fmt"
	"os"

	"github.com/mosherivkin/secgen/args"
	"github.com/mosherivkin/secgen/gen"
	"github.com/mosherivkin/secgen/help"
)

func main() {
	argsStr := os.Args[1:]
	isHelp, helpText := help.IsHelp(argsStr)
	if isHelp {
		fmt.Println(helpText)
		return
	}

	args, err := args.Parse(argsStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	secret := gen.Gen(*args)
	fmt.Println(secret)
}
