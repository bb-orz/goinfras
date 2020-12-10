package XEsOlivere

import (
	elasticv7 "github.com/olivere/elastic/v7"
)

type commonES struct {
	client *elasticv7.Client
}

// TODO 编写一些es client 的通用操作

func (c *commonES) Status() error {

	return nil
}
