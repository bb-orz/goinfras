package user

/*
数据访问层，实现具体数据持久化操作
直接返回error和执行结果
*/

type oauthDao struct{}

func NewOauthDao() *oauthDao {
	dao := new(oauthDao)
	return dao
}
