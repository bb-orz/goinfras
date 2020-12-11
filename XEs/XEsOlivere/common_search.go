package XEsOlivere

import (
	"context"
	elasticv7 "github.com/olivere/elastic/v7"
)

/*
一些常用的搜索方法
*/
type EsCommonSearch struct {
	client *elasticv7.Client
}

func (c *EsCommonSearch) Search(ctx context.Context, esIndex, esType string, query elasticv7.Query) (*elasticv7.SearchResult, error) {
	return c.client.Search().Index(esIndex).Type(esType).Query(query).Pretty(true).Human(true).Do(ctx)
}

func (c *EsCommonSearch) MultiSearch(ctx context.Context, esIndex string, requests ...*elasticv7.SearchRequest) (*elasticv7.MultiSearchResult, error) {
	return c.client.MultiSearch().Index(esIndex).Pretty(true).Human(true).Add(requests...).Do(ctx)
}

func (c *EsCommonSearch) SearchShards(ctx context.Context, indies ...string) (*elasticv7.SearchShardsResponse, error) {
	return c.client.SearchShards(indies...).Pretty(true).Human(true).Do(ctx)
}

/* 当数据大量分布在不同shards时，使用XPack Async Search 进行一步搜索 */

// func (c *EsCommonSearch) XPackAsyncSearchGet(docId string) {
// 	c.client.XPackAsyncSearchGet()
// }
//
// func (c *EsCommonSearch) XPackAsyncSearchSubmit() {
// 	c.client.XPackAsyncSearchSubmit()
// }
//
// func (c *EsCommonSearch) XPackAsyncSearchDelete() {
// 	c.client.XPackAsyncSearchDelete()
// }
