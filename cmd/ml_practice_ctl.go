package main

import (
	"fmt"

	"github.com/sheva0914/ml-practice/pkg/internal/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
