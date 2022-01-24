package main

import (
	"log"
	"os"
)

func main() {
	bData, err := os.ReadFile("./gargantua.sql")
	if err != nil {
		log.Fatal(err.Error())
	}

	tableParser := newTableParser()
	tables := tableParser.ParseString(string(bData))
	_ = tables

	print(genStruct(tables[0]))
}
