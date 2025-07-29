package html

import(
	"fmt"
	"os"
	"bufio"
)
// takes the html created and inserts it into the format of a typical html document, creates a new file for this called index.html
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
	file, err := os.Create("/home/elusama/projects/github.com/ECHIDNATHEG/GolangSSG/src/index.html")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(total_string)
	writer.Flush()
}