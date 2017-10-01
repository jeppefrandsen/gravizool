package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const gravizoBegin string = "http://g.gravizo.com/svg?"
const gravizoEnd string = "enduml"

var gravizoEncode = strings.NewReplacer(";", "%3B", " ", "%20", "\n", "%0A", "@", "%40",
	"(", "%28", ")", "%29", "*", "%2A", "\\", "%5C")
var gravizoDecode = strings.NewReplacer("%3B", ";", "%20", " ", "%0A", "\n", "%40", "@",
	"%2A", "*", "%5C", "\\")
var removeSemicolons = strings.NewReplacer(";", "")
var addSemicolons = strings.NewReplacer("\n", ";\n")

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convert(filename string, replacer *strings.Replacer, backup bool) {
	buffer, err := ioutil.ReadFile(filename)
	check(err)

	if backup {
		err = ioutil.WriteFile(fmt.Sprint(filename+".bak"), buffer, 0644)
		check(err)
	}

	text := string(buffer)

	for _, slice := range strings.Split(text, gravizoBegin) {
		if strings.Contains(slice, gravizoEnd) {
			subSlice := strings.Split(slice, gravizoEnd)
			if len(subSlice) > 0 {
				gravizoText := subSlice[0]
				convertedText := replacer.Replace(gravizoText)
				text = strings.Replace(text, gravizoText, convertedText, -1)
			}
		}
	}

	err = ioutil.WriteFile(filename, []byte(text), 0644)
	check(err)
}

func main() {
	encode := flag.String("e", "", "Encode the GitHub Markdown file")
	decode := flag.String("d", "", "Decode the GitHub Markdown file")
	fix := flag.String("f", "", "Fix the GitHub Markdown file")
	backup := flag.Bool("b", true, "Backup the GitHub Markdown file")

	flag.Parse()

	if len(*encode) > 0 {
		convert(*encode, gravizoEncode, *backup)
	} else if len(*decode) > 0 {
		convert(*decode, gravizoDecode, *backup)
	} else if len(*fix) > 0 {
		convert(*fix, removeSemicolons, *backup)
		convert(*fix, addSemicolons, *backup)
	}
}
