package XLogger

import (
	"context"
	"errors"
	"github.com/bb-orz/goinfras/XStore/XMongo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewMongoLogCore(cfg *Config) (zapcore.Core, error) {
	mongoWriter, err := NewMongoWriter(cfg)
	if err != nil {
		return nil, err
	}
	// 转成符合zapcore的输出接口类型
	mongoSyncer := zapcore.AddSync(mongoWriter)

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(defaultFormatConfig()),
		// 日志异步输出配置
		zapcore.NewMultiWriteSyncer(mongoSyncer),
		// 日志级别
		SettingLevelEnableFunc(cfg),
	), nil
}

func NewMongoWriter(cfg *Config) (*MongoWriter, error) {
	if XMongo.XClient() == nil {
		return nil, errors.New("XMongo Starter Not Setup! ")
	}

	writer := new(MongoWriter)
	writer.commonMongo = XMongo.XCommon(cfg.MongoLogDbName)
	writer.dbName = cfg.MongoLogDbName
	writer.collectionName = cfg.MongoLogCollection
	return writer, nil
}

// 实现io.Writer接口
type MongoWriter struct {
	commonMongo    *XMongo.CommonMongoDao
	dbName         string
	collectionName string
}

func (w *MongoWriter) Write(p []byte) (n int, err error) {
	_, err = w.commonMongo.InsertOne(context.TODO(), w.collectionName, p)
	if err != nil {
		XSyncError().Error("Mongo Sync Log Error!", zap.Error(err))
	}
	return 0, err
}
