package main

import (
	"fmt"
	"github.com/hfern/luao/lexer"
	"os"
)

func main() {
	file, err := os.Open("lua-test.lua")
	if err != nil {
		panic(err)
	}

	//src := "3*3"
	l := lexer.New(file)
	ch := l.Stream()

	for t := range ch {
		// Noone's interested in whitespace
		if t.Type() == lexer.Whitespace {
			continue
		}
		fmt.Println("Recived Tok:", lexer.TokenNames[t.Type()])
		fmt.Println("Line:", t.Line())
		fmt.Println("Bytes:", string(t.Bytes()))
		fmt.Println("")
	}

	fmt.Println("Done!")
}
