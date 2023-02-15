package main

import (
	"htmlink/linkparser"
	"htmlink/utils"
	"log"
)

const (
	example1Path = "resources/ex1.html"
	example2Path = "resources/ex2.html"
	example3Path = "resources/ex3.html"
	example4Path = "resources/ex4.html"
)

func main() {
	parser := linkparser.SimpleLinkParser()
	example1, err := utils.ReadFile(example1Path)
	example2, err := utils.ReadFile(example2Path)
	example3, err := utils.ReadFile(example3Path)
	example4, err := utils.ReadFile(example4Path)
	if err != nil {
		log.Fatal(err)
	}
	links1, err := parser.Parse(example1)
	links2, err := parser.Parse(example2)
	links3, err := parser.Parse(example3)
	links4, err := parser.Parse(example4)
	if err != nil {
		log.Fatal(err)
	}
	_ = links1
	_ = links2
	_ = links3
	_ = links4
}
