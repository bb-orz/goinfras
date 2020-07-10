package verified

import (
	"GoWebScaffold/infras/global"
	"GoWebScaffold/infras/mail"
	"GoWebScaffold/services"
	"fmt"
	"strconv"
)

/*Verified 校验领域层：实现邮件绑定、手机短信绑定相关的校验业务逻辑*/
type VerifiedDomain struct {
	cache *verifiedCache
}

func NewMailDomain() *VerifiedDomain {
	domain := new(VerifiedDomain)
	domain.cache = NewMailCache()
	return domain
}

// 生成邮箱验证码
func (domain *VerifiedDomain) genEmailVerifiedCode(uid int) (string, error) {
	// 生成6位随机字符串
	code := global.RandomString(6)

	// 保存到缓存
	err := domain.cache.SetUserVerifiedEmailCode(uid, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

// 构造邮件
func (domain *VerifiedDomain) sendEmail(email string, code string) error {
	from := "no-reply@" + global.Config().ServerName
	subject := "Verified Email Code From " + global.Config().AppName
	body := fmt.Sprintf("Verified Code: %s", code)
	return mail.SendSimpleMail(from, email, subject, body)
}

// 发送验证码到邮箱
func (domain *VerifiedDomain) SendEmail(dto services.SendEmailForVerifiedDTO) error {
	uid := int(dto.ID)
	email := dto.Email

	code, err := domain.genEmailVerifiedCode(uid)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendEmail(email, code)
}

// 验证邮箱
func (domain *VerifiedDomain) VerifiedEmail(uid int, vcode string) (bool, error) {

	// 缓存取出
	code, err := domain.cache.GetUserVerifiedEmailCode(uid)
	if err != nil {
		return false, err
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 生成手机短信验证码
func (domain *VerifiedDomain) genPhoneVerifiedCode(uid int) (string, error) {
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

// 构造短信
func (domain *VerifiedDomain) sendPhoneMsg(phone string, code string) error {

	return nil
}

// 发送验证码到手机短信
func (domain *VerifiedDomain) SendPhoneMsg(dto services.SendPhoneVerifiedCodeDTO) error {
	uid := int(dto.ID)
	phone := strconv.Itoa(int(dto.Phone))

	code, err := domain.genPhoneVerifiedCode(uid)
	if err != nil {
		return err
	}

	return domain.sendPhoneMsg(phone, code)
}

// 验证手机短信
func (domain *VerifiedDomain) VerifiedPhone(uid int, vcode string) (bool, error) {
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
