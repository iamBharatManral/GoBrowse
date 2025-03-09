package parser

import "fmt"

type HTMLParser struct{}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

func (p *HTMLParser) Parse(body string) {
	fmt.Println(body)
}
