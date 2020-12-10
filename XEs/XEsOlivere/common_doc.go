package XEsOlivere

import (
	"context"
	"fmt"
	elasticv7 "github.com/olivere/elastic/v7"
	"strconv"
)

/*一些常用的文档操作方法*/
type EsCommonDoc struct {
	client *elasticv7.Client
}

// 查看某文档是否存在,给定文档ID查询
func (c *EsCommonDoc) IsDocExists(id int, index string) (bool, error) {
	var err error
	exist, err := c.client.Exists().Index(index).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return true, nil
}

// 获取文档
func (c *EsCommonDoc) GetDoc(id int, index string) (*elasticv7.GetResult, error) {
	esResponse, err := c.client.Get().Index(index).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return esResponse, nil
}

// 添加文档
func (c *EsCommonDoc) AddDoc(id int, doc string, index string) (*elasticv7.IndexResponse, error) {
	rsp, err := c.client.Index().Index(index).Id(strconv.Itoa(id)).BodyJson(doc).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// 更新文档
func (c *EsCommonDoc) UpdateDoc(updateField *map[string]interface{}, id int, index string) (*elasticv7.UpdateResponse, error) {
	rsp, err := c.client.Update().Index(index).Id(strconv.Itoa(id)).Doc(updateField).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rsp, nil
}

// 删除文档
func (c *EsCommonDoc) DeleteDoc(id int, index string) (*elasticv7.DeleteResponse, error) {
	rsp, err := c.client.Delete().Index(index).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
