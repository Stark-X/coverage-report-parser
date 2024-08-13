package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"coverage-report-parser/parser"
	"coverage-report-parser/parser/lcov"
	"coverage-report-parser/reader"
)

var (
	reportFilePath *string
)

func main() {
	reportFilePath = flag.String("report", "", "The report file to parse")
	flag.Parse()
	fmt.Println("Report file:", *reportFilePath)

	r, err := reader.GetReader(*reportFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer func() {
		if f := r.(*os.File); f != nil {
			f.Close()
		}
		if r := recover(); r != nil {
			fmt.Println("Recovered in main: ", r)
		}
	}()

	var parser parser.Parser = lcov.NewParser(r)
	res, err := parser.Parse()
	if err != nil {
		log.Fatalf("Error parsing report: %v", err)
		return
	}

	fmt.Println(res)
}
