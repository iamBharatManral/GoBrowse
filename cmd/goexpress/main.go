package main

import (
	connect "github.com/iamBharatManral/GoBrowse/internal/client"
	"github.com/iamBharatManral/GoBrowse/internal/parser"
)

func main() {
	client := connect.Http()
	_ = client.Get("https://example.org")

	htmlParser := parser.NewHTMLParser()
	htmlParser.Parse(string(client.ResBody))
}
