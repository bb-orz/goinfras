package XEsOlivere

import (
	elasticv7 "github.com/olivere/elastic/v7"
)

var esClient *elasticv7.Client

func XClient() *elasticv7.Client {
	return esClient
}

func XFClient(f func(c *elasticv7.Client) error) error {
	return f(esClient)
}

func XCommon() *EsCommon {
	c := new(EsCommon)
	c.client = XClient()
	return c
}