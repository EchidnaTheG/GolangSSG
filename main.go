package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const TESTFILE = "/home/elusama/projects/github.com/ECHIDNATHEG/GolangSSG/Test.md"

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
	var multiline bool
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

		temp, multiline = HandleText(Line)
		if multiline {

			if multiline && !counter {
				counter = true
				Line = fmt.Sprintf("<li>\n  %v\n", temp.Form)
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
			Line = fmt.Sprintf("</li>\n\n%v", temp.Form)
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

func HandleText(line string) (TEXT, bool) {
	original := line
	if strings.TrimSpace(line) == "" {
		// Line is empty or contains only whitespace
		return TEXT{Form: line, hasChildren: false}, false
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
		line = re.ReplaceAllString(line, "<ul>$1</ul>")
		return TEXT{Form: line, hasChildren: hasChanges}, true
	}
	re = regexp.MustCompile(`^\d+\.\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, "<ol>$1</ol>")
		return TEXT{Form: line, hasChildren: hasChanges}, true
	}

	line = fmt.Sprintf("<p>%v</p>", line)
	return TEXT{Form: line, hasChildren: hasChanges}, false
}

func main() {
	total_string := ScanMarkdownFile(TESTFILE)
	fmt.Println(total_string)
}
