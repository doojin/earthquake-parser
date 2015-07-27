package main

import (
	"github.com/doojin/earthquake-parser/parser"
	"fmt"
)

func main() {
	result := parser.SendRequest("2014-01-01", "2014-01-02", 1, 10)
	fmt.Println(result)
}
