package XEsOfficial

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
)

type EsCommon struct {
	client *esv8.Client
}

// TODO 编写一些es client 的通用操作

func (c *EsCommon) Status() error {

	return nil
}
