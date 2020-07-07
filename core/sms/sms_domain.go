package sms

import (
	"GoWebScaffold/infras/global"
	"GoWebScaffold/services"
	"strconv"
)

/*领域层：实现具体业务逻辑*/
type SmsDomain struct {
	cache *smsCache
}

func NewMailDomain() *SmsDomain {
	domain := new(SmsDomain)
	domain.cache = NewSmsCache()
	return domain
}

// 生成验证码
func (domain *SmsDomain) genPhoneVerifiedCode(uid int) (string, error) {
	var err error
	var code string
	// 生成4位随机数字
	code, err = global.RandomNumber(4)
	if err != nil {
		return "", nil
	}

	// 保存到缓存
	err = domain.cache.SetUserVerifiedPhoneCode(uid, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (domain *SmsDomain) sendPhoneMsg(phone string, code string) error {

	return nil
}

// 发送验证码到手机短信
func (domain *SmsDomain) SendPhoneMsg(dto services.SendPhoneVerifiedCodeDTO) error {
	uid := int(dto.ID)
	phone := strconv.Itoa(int(dto.Phone))

	code, err := domain.genPhoneVerifiedCode(uid)
	if err != nil {
		return err
	}

	return domain.sendPhoneMsg(phone, code)
}

// 验证邮箱
func (domain *SmsDomain) VerifiedPhone(uid int, vcode string) (bool, error) {
	// 缓存取出
	code, err := domain.cache.GetUserVerifiedPhoneCode(uid)
	if err != nil {
		return false, err
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}
