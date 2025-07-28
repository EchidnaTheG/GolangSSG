package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const TESTFILE = "/home/elusama/projects/github.com/ECHIDNATHEG/GolangSSG/Test2.md"

type TEXT struct {
	Form        string
	Sliced      []string
	hasChildren bool
}

type HEADER struct {
	Form string
}

func ScanMarkdownFile(filename string) string {
	file, err := os.Open(filename)
	
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	defer file.Close()
	var multiline bool
	var class string
	var prevclass string
	counter := false
	var temp TEXT
	totalString := ""

	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		Line := Scanner.Text()
		if val, ok := ParseHeaders(Line); ok {
			Line = val.Form + "\n"
			totalString += Line
			continue
		}

		temp, multiline,class = HandleText(Line)
		if multiline {
			prevclass = class
			if multiline && !counter {
				counter = true
				Line = fmt.Sprintf("<%v>\n  %v\n",class, temp.Form)
				totalString += Line
				continue

			}
			if multiline && counter {
				Line = fmt.Sprintf("  %v\n", temp.Form)
				totalString += Line
				continue
			}

		}
		if counter && !multiline {
			counter = false
			Line = fmt.Sprintf("</%v>\n\n%v",prevclass, temp.Form)
			totalString += Line
			continue
		}
		Line = temp.Form + "\n"
		totalString += Line
	}

	return totalString
}

func ParseHeaders(line string) (HEADER, bool) {
	re := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		level := len(match[1])
		header := HEADER{fmt.Sprintf("<h%d>%s</h%d>", level, match[2], level)}
		return header, true
	}
	return HEADER{}, false
}

func HandleText(line string) (TEXT, bool, string) {
	original := line
	if strings.TrimSpace(line) == "" {
		// Line is empty or contains only whitespace
		return TEXT{Form: line, hasChildren: false}, false,""
	}
	// Bold: **text** -> <strong>text</strong>
	re := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	line = re.ReplaceAllString(line, "<strong>$1</strong>")

	// Italic: *text* -> <em>text</em>
	re = regexp.MustCompile(`\*([^*]+)\*`)
	line = re.ReplaceAllString(line, "<em>$1</em>")

	// Strikethrough: ~~text~~ -> <del>text</del>
	re = regexp.MustCompile(`~~([^~]+)~~`)
	line = re.ReplaceAllString(line, "<del>$1</del>")

	// Inline code: `text` -> <code>text</code>
	re = regexp.MustCompile("`([^`]+)`")
	line = re.ReplaceAllString(line, "<code>$1</code>")
	hasChanges := original != line

	re = regexp.MustCompile(`^-\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, "<li>$1</li>")
		return TEXT{Form: line, hasChildren: hasChanges}, true,"ul"
	}
	re = regexp.MustCompile(`^\d+\.\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, "<li>$1</li>")
		return TEXT{Form: line, hasChildren: hasChanges}, true,"ol"
	}

	line = fmt.Sprintf("<p>%v</p>", line)
	return TEXT{Form: line, hasChildren: hasChanges}, false, ""
}



func WriteToHTML(total_string string){
	total_string = fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    %v
</body>
</html>`, total_string)
	file, err := os.Create("/home/elusama/projects/github.com/ECHIDNATHEG/GolangSSG/index.html")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(total_string)
	writer.Flush()
}

func main() {
	total_string := ScanMarkdownFile(TESTFILE)
	fmt.Println(total_string)
	WriteToHTML(total_string)
}
