package XEsOlivere

import (
	elasticv7 "github.com/olivere/elastic/v7"
)

/*一些常用的索引库操作方法*/

type EsCommonIndex struct {
	client *elasticv7.Client
}

func (c *EsCommonIndex) IndexExists(indices ...string) {

	// exists := c.client.IndexExists(indices...)
	// b, e := exists.Do(context.Background())
}
