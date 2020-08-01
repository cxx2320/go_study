package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

var t, token xml.Token
var err error

func main() {
	XmlTo()
}

func XmlTo() {
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	token, _ := p.Token()
	fmt.Println(token)
	fmt.Printf("%s\n", p)
}
