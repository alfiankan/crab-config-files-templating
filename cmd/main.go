package main

import (
	"config-replacer/replacer"
	"fmt"
)

func main() {
	if err := replacer.RootCLI().Execute(); err != nil {
		fmt.Println("Something went wrong", err.Error())
	}
}
