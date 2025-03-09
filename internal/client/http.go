package connect

import (
	"io"
	"log"
	"net/http"
)

type HttpClient struct {
	ResBody  string
	Response *http.Response
	Request  *http.Request
}

func Http() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) Get(url string) error {
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	c.ResBody = string(body)
	c.Request = res.Request
	c.Response = res
	return err
}
