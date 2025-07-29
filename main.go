package main

import (
	"fmt"
	"github.com/EchidnaTheG/GolangSSG/internal/file"
	"github.com/EchidnaTheG/GolangSSG/internal/html"

)

const TESTFILE = "/home/elusama/projects/github.com/ECHIDNATHEG/GolangSSG/public/Test.md"

func main() {
	total_string := file.ScanMarkdownFile(TESTFILE)
	fmt.Println(total_string)
	html.WriteToHTML(total_string)
}
