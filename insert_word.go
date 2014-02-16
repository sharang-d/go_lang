package main

import (
	"io/ioutil"
	"regexp"
	"os"
)

func main() {
	path := "input.txt"
	data, _ := ioutil.ReadFile(path)
	text := string(data)
	word, replacement := "word", "inserted word"
	re := regexp.MustCompile("(\\b)" + word + "(\\b)")
	text = re.ReplaceAllString(text, "${1}" + replacement + "${2}")
	file, _ := os.Create(path)
	file.WriteString(text)
	file.Sync()
	file.Close()
}
