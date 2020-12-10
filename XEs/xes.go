package XEs

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
)

var esClient *esv8.Client

func XClient() *esv8.Client {
	return esClient
}

func XFClient(f func(c *esv8.Client) error) error {
	return f(esClient)
}

func XCommonES() *commonES {
	c := new(commonES)
	c.client = XClient()
	return c
}
