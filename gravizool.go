package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const gravizoBegin string = "(http://g.gravizo.com/svg?"

//const gravizoEnd string = "enduml)"

var encoder = strings.NewReplacer(";", "%3B", " ", "%20", "\n", "%0A", "@", "%40",
	"(", "%28", ")", "%29", "*", "%2A", "\\", "%5C")
var decoder = strings.NewReplacer("%3B", ";", "%20", " ", "%0A", "\n", "%40", "@",
	"%2A", "*", "%5C", "\\")

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func findMatchingClose(text string, opening rune, closing rune) int {
	openingCount := 1
	for i, ch := range text {
		if ch == opening {
			openingCount++
		} else if ch == closing {
			openingCount--
			if openingCount == 0 {
				return i
			}
		}
	}
	return 0
}

func convert(filename string, replacer *strings.Replacer, backup bool) {
	buffer, err := ioutil.ReadFile(filename)
	check(err)

	if backup {
		err = ioutil.WriteFile(fmt.Sprint(filename+".bak"), buffer, 0644)
		check(err)
	}

	text := string(buffer)

	for offset, slice := range strings.Split(text, gravizoBegin) {
		if offset == 0 {
			continue
		}
		closeOffset := findMatchingClose(slice, '(', ')')
		if closeOffset > 0 {
			gravizoText := slice[:closeOffset]
			if len(gravizoText) > 0 {
				convertedText := replacer.Replace(gravizoText)
				text = strings.Replace(text, gravizoText, convertedText, -1)
			}
		}
	}

	err = ioutil.WriteFile(filename, []byte(text), 0644)
	check(err)
}

func main() {
	encode := flag.String("e", "", "Encode the given GitHub Markdown file")
	decode := flag.String("d", "", "Decode the given GitHub Markdown file")
	backup := flag.Bool("b", true, "Backup GitHub Markdown file before encode/decode")

	flag.Parse()

	if len(*encode) > 0 {
		convert(*encode, encoder, *backup)
	} else if len(*decode) > 0 {
		convert(*decode, decoder, *backup)
	}
}
