package user

/*
数据访问层，实现具体数据持久化操作
直接返回error和执行结果
*/

type oauthDAO struct{}

func NewOauthDAO() *oauthDAO {
	dao := new(oauthDAO)
	return dao
}

// 通过openid unionid 获取oauth user 信息
func (d *oauthDAO) GetOauthInfoByOpenId(openId, unionId string) (*UserOauth, error) {

	return nil, nil
}
