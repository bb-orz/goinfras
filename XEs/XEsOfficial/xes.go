package XEsOfficial

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
)

func XClient() *esv8.Client {
	return esClient
}

func XFClient(f func(c *esv8.Client) error) error {
	return f(esClient)
}

func XCommonES() *EsCommon {
	c := new(EsCommon)
	c.client = XClient()
	return c
}
