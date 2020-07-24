package main

/* stripansi.go - Strips ANSI Colours for content provided in input. Based on
   work of acarl005
*/
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

// Strip - Strip ANSi colours out from specified str via regex
func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}

func main() {

	// Read input from user
	sc := bufio.NewScanner(os.Stdin)

	// Read line from input
	for sc.Scan() {

		// Strips the colors out fo each line read
		line := Strip(sc.Text())

		// Print out the line read
		fmt.Println(line)
	}
}
