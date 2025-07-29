package file
import(
	"os"
	"fmt"
	"bufio"
	"github.com/EchidnaTheG/GolangSSG/internal/types"
	"github.com/EchidnaTheG/GolangSSG/internal/parser"
)


func ScanMarkdownFile(filename string) string {
	//IO stuff, getting md file
	file, err := os.Open(filename)
	
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	defer file.Close()

	//variables for state management in parsing loop
	var multiline bool
	var class string
	var prevclass string
	IsCodeBlockDone := true
	counter := false
	var temp types.TEXT
	totalString := ""

	//parsing loop with scanner
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {

		Line := Scanner.Text()

		//checks if is start of codeblock, since IsCodeBlock starts as true, it can be set to false once there is a code block in progress
		if val, ok := parser.IsCodeBlock(Line); ok && IsCodeBlockDone{
			Line = fmt.Sprintf(`<pre>`+"\n" +` <code class="%v">`+"\n ",val)
			totalString += Line
			IsCodeBlockDone = false
			continue
		}
		//what runs when there is a code block interior building
		if !IsCodeBlockDone{
			//checks if codeBlock is done
			if val, ok := parser.IsCodeBlockDoneCheck(Line); ok{
				IsCodeBlockDone = true
				Line = val +"\n"
				totalString += Line
				continue
			}
			//builds the interior of the codeblock if not done and repeats
			totalString += Line +"\n"
			continue
		}
		//parsers headers like h1 and etc
		if val, ok := parser.ParseHeaders(Line); ok {
			Line = val.Form + "\n"
			totalString += Line
			continue
		}
		//handles most of the text in the md file, treats also lists which is what multiline is for
		temp, multiline,class = parser.HandleText(Line)
		if multiline {
			prevclass = class
			//counter is false while not multiline, its true while multiline build is in progress, this is the init part
			if multiline && !counter {
				counter = true
				Line = fmt.Sprintf("<%v>\n  %v\n",class, temp.Form)
				totalString += Line
				continue

			}
			//this is the middle of the multiline, while its still multiline and still building
			if multiline && counter {
				Line = fmt.Sprintf("  %v\n", temp.Form)
				totalString += Line
				continue
			}

		}
		//this is what runs when it stops being multiline and counter is still true, sets counter back to false
		if counter && !multiline {
			counter = false
			Line = fmt.Sprintf("</%v>\n\n%v",prevclass, temp.Form)
			totalString += Line
			continue
		}
		//this is what runs if its a normal non multiline element
		Line = temp.Form + "\n"
		totalString += Line
	}
	//complete html file
	return totalString
}