package parser
// takes in line, returns header struct, the form field specifically has the html 

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/EchidnaTheG/GolangSSG/internal/types"
)


func ParseHeaders(line string) (types.HEADER, bool) {
	re := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		level := len(match[1])
		text := applyTextFormatting(match[2])
		header := types.HEADER{Form: fmt.Sprintf("<h%d>%s</h%d>", level, text, level)}
		return header, true
	}
	return types.HEADER{}, false
}

// Hfunction for formatting md into their html equivalents, can be used everywhere where neeeded
func applyTextFormatting(line string) string {
    // Bold and Italic: ***text*** -> <em><strong>text</strong></em>
    re := regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)
    line = re.ReplaceAllString(line, "<em><strong>$1</strong></em>")

    // Bold: **text** -> <strong>text</strong>
    re = regexp.MustCompile(`\*\*(.*?)\*\*`)
    line = re.ReplaceAllString(line, "<strong>$1</strong>")

    // Italic: *text* -> <em>text</em>
    re = regexp.MustCompile(`\*(.*?)\*`)
    line = re.ReplaceAllString(line, "<em>$1</em>")

    // Strikethrough: ~~text~~ -> <del>text</del>
    re = regexp.MustCompile(`~~(.*?)~~`)
    line = re.ReplaceAllString(line, "<del>$1</del>")

    // Inline code: `text` -> <code>text</code>
    re = regexp.MustCompile("`([^`]+)`")
    line = re.ReplaceAllString(line, "<code>$1</code>")
    
    return line
}

//logic for handling the majority of md cases, does both dealing with multiline and single line
func HandleText(line string) (types.TEXT, bool, string) {
	original := line
	if strings.TrimSpace(line) == "" {
		// Line is empty or contains only whitespace
		return types.TEXT{Form: line, HasChildren: false}, false,""
	}
	line = applyTextFormatting(line)

	hasChanges := original != line

	
	re := regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, `<img src="$2" alt="$1">`)
		
	}

	re = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`) // [text](url) or [text](url "title")
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, `<a href="$2">$1</a>`)
	}
	
	
	re = regexp.MustCompile(`^-\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, "<li>$1</li>")
		return types.TEXT{Form: line, HasChildren: hasChanges}, true,"ul"
	}

	

	re = regexp.MustCompile(`^\d+\.\s+(.+)$`)
	if match := re.FindStringSubmatch(line); match != nil {
		line = re.ReplaceAllString(line, "<li>$1</li>")
		return types.TEXT{Form: line, HasChildren: hasChanges}, true,"ol"
	}

	line = fmt.Sprintf("<p>%v</p>", line)
	return types.TEXT{Form: line, HasChildren: hasChanges}, false, ""
}
// confirms the start of a codeblock
func IsCodeBlock(line string) (string, bool){
  re := regexp.MustCompile(`^` + "```" + `(\w*)$`)
    if match := re.FindStringSubmatch(line); match != nil {
        return  match[1] ,true // Returns the language name
    }
    return  "", false
}
//handles code block in process and returns when its done
func IsCodeBlockDoneCheck(line string) (string,bool){
	 re := regexp.MustCompile("^```$")
	 if match := re.FindStringSubmatch(line); match !=nil{
		rep := fmt.Sprintf(` </code>`+"\n" +`</pre>`)
		line = re.ReplaceAllString(line, rep)
		return line, true
	 }
	 return "",false
}