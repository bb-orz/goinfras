package mongoStore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommonMongoDao struct {
	client    *mongo.Client
	defaultDb *mongo.Database
}

func NewCommonMongoDao(dbName string) *CommonMongoDao {
	c := new(CommonMongoDao)
	c.client = Client()
	c.defaultDb = c.client.Database(dbName)
	return c
}

func (mp *CommonMongoDao) M(colName string, f func(c *mongo.Collection) error) error {
	collection := mp.defaultDb.Collection(colName)
	return f(collection)
}

func (mp *CommonMongoDao) DM(dbName, colName string, f func(c *mongo.Collection) error) error {
	collection := mp.client.Database(dbName).Collection(colName)
	return f(collection)
}

/*
通用新增数据到mongo collection
*/
func (mp *CommonMongoDao) InsertOne(ctx context.Context, collectionName string, document interface{}, opts ...*options.InsertOneOptions) (insertID interface{}, err error) {
	one, err := mp.defaultDb.Collection(collectionName).InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}
	return one.InsertedID, nil
}

func (mp *CommonMongoDao) InsertMany(ctx context.Context, collectionName string, documents []interface{}, opts ...*options.InsertManyOptions) (insertIDs []interface{}, err error) {
	many, err := mp.defaultDb.Collection(collectionName).InsertMany(ctx, documents, opts...)
	if err != nil {
		return nil, err
	}
	return many.InsertedIDs, nil
}

/*
通用更新数据

*/
func (mp *CommonMongoDao) UpdateOne(ctx context.Context, collectionName string, filter, updater interface{}, opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return mp.defaultDb.Collection(collectionName).UpdateOne(ctx, filter, updater, opts...)
}

func (mp *CommonMongoDao) UpdateMany(ctx context.Context, collectionName string, filter, updater interface{}, opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return mp.defaultDb.Collection(collectionName).UpdateMany(ctx, filter, updater, opts...)
}

func (mp *CommonMongoDao) ReplaceOne(ctx context.Context, collectionName string, filter, replacement interface{}, opts ...*options.ReplaceOptions) (result *mongo.UpdateResult, err error) {
	return mp.defaultDb.Collection(collectionName).ReplaceOne(ctx, filter, replacement, opts...)
}

/*
通用删除数据

*/
func (mp *CommonMongoDao) DeleteOne(ctx context.Context, collectionName string, filter interface{}, opts ...*options.DeleteOptions) (deleteCount int64, err error) {
	deleteResult, err := mp.defaultDb.Collection(collectionName).DeleteOne(ctx, filter, opts...)
	if err != nil {
		return -1, err
	}
	return deleteResult.DeletedCount, nil
}

func (mp *CommonMongoDao) DeleteMany(ctx context.Context, collectionName string, filter interface{}, opts ...*options.DeleteOptions) (deleteCount int64, err error) {
	deleteResult, err := mp.defaultDb.Collection(collectionName).DeleteMany(ctx, filter, opts...)
	if err != nil {
		return -1, err
	}
	return deleteResult.DeletedCount, nil
}

type FindResult interface{}
type FindResults []interface{}

/*
查找文档
*/
func (mp *CommonMongoDao) Find(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOptions) (cursor *mongo.Cursor, err error) {
	return mp.defaultDb.Collection(collectionName).Find(ctx, filter, opts...)
}

func (mp *CommonMongoDao) FindOne(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return mp.defaultDb.Collection(collectionName).FindOne(ctx, filter, opts...)
}

func (mp *CommonMongoDao) FindOneAndDelete(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOneAndDeleteOptions) (singleResult *mongo.SingleResult) {
	return mp.defaultDb.Collection(collectionName).FindOneAndDelete(ctx, filter, opts...)
}

func (mp *CommonMongoDao) FindOneAndUpdate(ctx context.Context, collectionName string, filter, updater interface{}, opts ...*options.FindOneAndUpdateOptions) (result *mongo.SingleResult) {
	return mp.defaultDb.Collection(collectionName).FindOneAndUpdate(ctx, filter, updater, opts...)
}

func (mp *CommonMongoDao) FindOneAndReplace(ctx context.Context, collectionName string, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) (singleResult *mongo.SingleResult) {
	return mp.defaultDb.Collection(collectionName).FindOneAndReplace(ctx, filter, replacement, opts...)
}

/*
批量操作
*/
func (mp *CommonMongoDao) BulkWrite(ctx context.Context, collectionName string, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (bulkResult *mongo.BulkWriteResult, err error) {
	return mp.defaultDb.Collection(collectionName).BulkWrite(ctx, models, opts...)
}

// 查找文档数
func (mp *CommonMongoDao) CountDocuments(ctx context.Context, collectionName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mp.defaultDb.Collection(collectionName).CountDocuments(ctx, filter, opts...)
}

/*
聚合计算
*/
func (mp *CommonMongoDao) Aggregate(ctx context.Context, collectionName string, pipline interface{}, opts ...*options.AggregateOptions) (cursor *mongo.Cursor, err error) {
	return mp.defaultDb.Collection(collectionName).Aggregate(ctx, pipline, opts...)
}
