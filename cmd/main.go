package main

import (
	"fmt"

	"github.com/alfiankan/crab-config-files-templating/replacer"
)

func main() {
	if err := replacer.RootCLI().Execute(); err != nil {
		fmt.Println("Something went wrong", err.Error())
	}
}
