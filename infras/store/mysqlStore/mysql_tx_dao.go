package mysqlStore

import (
	"context"
	"database/sql"
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

// 创建一个事务
func (m *BaseDao) NewTx(ctx context.Context, options *sql.TxOptions) (*MysqlTx, error) {
	var err error
	mysqlTx := new(MysqlTx)
	mysqlTx.tx, err = m.db.BeginTx(ctx, options)
	if err != nil {
		return nil, err
	}
	return mysqlTx, nil
}

type MysqlTx struct {
	tx *sql.Tx
}

// 以下提供通用的几个基于事务的curd方法，具体构建方式可查看本目录的README
/*
根据条件获取单条数据
@param tableName 	string					查询的表名
@param where 		map[string]interface{}	查询条件
@param selectField 	[]string				查询选择返回的字段
@param result 		DaoMysqlSchema			带表结构存储结果的指针，接收返回的数据，实现DaoMysqlSchema接口
*/
func (mtx *MysqlTx) GetOne(tableName string, where map[string]interface{}, selectField []string, result interface{}) error {
	if mtx.tx == nil {
		return errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildSelect(tableName, where, selectField)
	if err != nil {
		mtx.tx.Rollback()
		return err
	}

	row, err := mtx.tx.Query(condition, values...)
	if err != nil {
		return err
	}
	defer row.Close()
	return scanner.Scan(row, result)
}

/*
根据条件获取多条数据
@param tableName 	string					查询的表名
@param where 		map[string]interface{}	查询条件
@param selectField 	[]string				查询选择返回的字段
@param results 		[]DaoMysqlSchema		带表结构存储结果的指针数组，接收返回的数据，实现DaoMysqlSchema接口
*/
func (mtx *MysqlTx) GetMulti(tableName string, where map[string]interface{}, selectField []string, results []interface{}) error {
	if mtx.tx == nil {
		return errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildSelect(tableName, where, selectField)
	if err != nil {
		return err
	}

	rows, err := mtx.tx.Query(condition, values...)
	if err != nil || nil == rows {
		return err
	}
	defer rows.Close()
	return scanner.Scan(rows, results)
}

/*
插入数据
@param tableName 	string						表名
@param data 		[]map[string]interface{}	插入数据

@return LastInsertId int64						返回最新的插入id
*/
func (mtx *MysqlTx) Insert(tableName string, data []map[string]interface{}) (int64, error) {
	if mtx.tx == nil {
		return -1, errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildInsert(tableName, data)
	if err != nil {
		return -1, err
	}
	result, err := mtx.tx.Exec(condition, values...)
	if err != nil || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

/*
插入数据,已存在则忽略
@param tableName 	string						表名
@param data 		[]map[string]interface{}	插入数据

@return LastInsertId int64						返回最新的插入id
*/
func (mtx *MysqlTx) InsertIgnore(tableName string, data []map[string]interface{}) (int64, error) {
	if mtx.tx == nil {
		return -1, errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildInsertIgnore(tableName, data)
	if err != nil {
		return -1, err
	}
	result, err := mtx.tx.Exec(condition, values...)
	if err != nil || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

/*
插入数据,已存在则替换
@param tableName 	string						表名
@param data 		[]map[string]interface{}	插入数据

@return LastInsertId int64						返回最新的插入id
*/
func (mtx *MysqlTx) InsertReplace(tableName string, data []map[string]interface{}) (int64, error) {
	if mtx.tx == nil {
		return -1, errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildReplaceInsert(tableName, data)
	if err != nil {
		return -1, err
	}
	result, err := mtx.tx.Exec(condition, values...)
	if err != nil || nil == result {
		return -1, err
	}
	return result.LastInsertId()
}

/*
更新数据
@param tableName 	string						表名
@param where 		map[string]interface{}		查询条件
@param data 		[]map[string]interface{}	更新数据

@return RowsAffected int64						更新影响的行数
*/
func (mtx *MysqlTx) Update(tableName string, where, data map[string]interface{}) (int64, error) {
	if mtx.tx == nil {
		return -1, errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildUpdate(tableName, where, data)
	if err != nil {
		return -1, err
	}
	result, err := mtx.tx.Exec(condition, values...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

/*
删除数据
@param tableName 	string						表名
@param where 		map[string]interface{}		查询条件

@return RowsAffected int64						删除影响的行数
*/func (mtx *MysqlTx) Delete(tableName string, where map[string]interface{}) (int64, error) {
	if mtx.tx == nil {
		return -1, errors.New("sql.Tx pointer couldn't be nil")
	}
	condition, values, err := builder.BuildDelete(tableName, where)
	if err != nil {
		return -1, err
	}
	result, err := mtx.tx.Exec(condition, values...)
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
