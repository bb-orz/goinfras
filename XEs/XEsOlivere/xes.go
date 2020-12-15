package XEsOlivere

import (
	elasticv7 "github.com/olivere/elastic/v7"
)

func XClient() *elasticv7.Client {
	return esClient
}

func XFClient(f func(c *elasticv7.Client) error) error {
	return f(esClient)
}

func XCommonDoc() *EsCommonDoc {
	c := new(EsCommonDoc)
	c.client = XClient()
	return c
}

func XCommonSearch() *EsCommonSearch {
	c := new(EsCommonSearch)
	c.client = XClient()
	return c
}
