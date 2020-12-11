package XEsOlivere

import (
	"context"
	"errors"
	"fmt"
	elasticv7 "github.com/olivere/elastic/v7"
	"strconv"
)

/*一些常用的文档操作方法*/
type EsCommonDoc struct {
	client *elasticv7.Client
}

// 查看某文档是否存在,给定文档ID查询
func (c *EsCommonDoc) IsDocExists(docId, esType, esIndex string) (bool, error) {
	var err error
	exist, err := c.client.Exists().Index(esIndex).Type(esType).Id(docId).Do(context.Background())
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return true, nil
}

// 获取文档
func (c *EsCommonDoc) GetDoc(docId, esIndex string) (*elasticv7.GetResult, error) {
	esResponse, err := c.client.Get().Index(esIndex).Id(docId).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return esResponse, nil
}

// 添加文档
func (c *EsCommonDoc) AddDoc(docId, esType, esIndex string, doc interface{}) (*elasticv7.IndexResponse, error) {
	var err error
	var isExsit bool
	var rsp *elasticv7.IndexResponse
	isExsit, err = c.IsDocExists(docId, esType, esIndex)
	if err != nil {
		return nil, err
	}

	if !isExsit {
		rsp, err = c.client.Index().Index(esIndex).Type(esType).Id(docId).BodyJson(doc).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return rsp, nil
}

// 批量插入
func (c *EsCommonDoc) BatchAddDoc(esIndex, esType string, datas ...interface{}) error {

	bulkService := elasticv7.NewBulkService(c.client)
	for i, data := range datas {
		doc := elasticv7.NewBulkIndexRequest().Index(esIndex).Type(esType).Id(strconv.Itoa(i)).Doc(data)
		bulkService = bulkService.Add(doc)
	}

	response, err := bulkService.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	failed := response.Failed()
	iter := len(failed)
	if response.Errors {
		return errors.New(fmt.Sprintf("error: %v, %v\n", response.Errors, iter))
	}
	return nil
}

// 更新文档
func (c *EsCommonDoc) UpdateDoc(updateField *map[string]interface{}, docId, esType, esIndex string) (*elasticv7.UpdateResponse, error) {
	var err error
	var isExsit bool
	var rsp *elasticv7.UpdateResponse
	isExsit, err = c.IsDocExists(docId, esType, esIndex)
	if err != nil {
		return nil, err
	}

	if !isExsit {
		rsp, err = c.client.Update().Index(esIndex).Type(esType).Id(docId).Doc(updateField).Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return rsp, nil
}

// 删除文档
func (c *EsCommonDoc) DeleteDoc(docId, esType, esIndex string) (*elasticv7.DeleteResponse, error) {
	var err error
	var isExsit bool
	var rsp *elasticv7.DeleteResponse
	isExsit, err = c.IsDocExists(docId, esType, esIndex)
	if err != nil {
		return nil, err
	}

	if !isExsit {
		rsp, err = c.client.Delete().Index(esIndex).Type(esType).Id(docId).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}
