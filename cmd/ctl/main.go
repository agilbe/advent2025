package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

const solverTemplate = `package internal

func SolveDay{{.Day}}(input string) string {
    // TODO: Implement Day {{.Day}}
    return ""
}

func init() {
    Register({{.Day}}, SolveDay{{.Day}})
}
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ctl <day-number>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day:", err)
		return
	}

	createDay(day)
}

func createDay(day int) {
	dayPadded := fmt.Sprintf("%02d", day)

	solverDir := filepath.Join("internal")
	if err := os.MkdirAll(solverDir, 0755); err != nil {
		panic(err)
	}

	solverPath := filepath.Join(solverDir, "day"+dayPadded+".go")
	if _, err := os.Stat(solverPath); err == nil {
		fmt.Println("Solver already exists:", solverPath)
	} else {
		f, err := os.Create(solverPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		t := template.Must(template.New("solver").Parse(solverTemplate))
		t.Execute(f, map[string]int{"Day": day})
		fmt.Println("Created solver:", solverPath)
	}

	inputDir := filepath.Join("inputs")
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		panic(err)
	}

	inputPath := filepath.Join(inputDir, "day"+dayPadded+".txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		os.WriteFile(inputPath, []byte(""), 0644)
		fmt.Println("Created input file:", inputPath)
	} else {
		fmt.Println("Input file already exists:", inputPath)
	}
}
