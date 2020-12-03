package verified

import (
	"GoWebScaffold/core"
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
func (domain *VerifiedDomain) genEmailVerifiedCode(uid uint) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code = XGlobal.RandomString(6)

	// 保存到缓存
	err = domain.cache.SetUserVerifiedEmailCode(uid, code)
	if err != nil {
		return "", core.WrapError(err, core.ErrorFormatDomainCacheSet, DomainName, "SetUserVerifiedEmailCode")
	}

	return code, nil
}

// 构造验证邮箱邮件
func (domain *VerifiedDomain) sendValidateEmail(email string, code string) error {
	from := "no-reply@" + XGlobal.Config().ServerName
	subject := "Verified Email Code From " + XGlobal.Config().AppName
	body := fmt.Sprintf("Verified Code: %s", code)
	return XMail.SendSimpleMail(from, email, subject, body)
}

// 发送验证码到邮箱
func (domain *VerifiedDomain) SendValidateEmail(dto services.SendEmailForVerifiedDTO) error {
	var err error
	var code string

	uid := dto.ID
	email := dto.Email

	code, err = domain.genEmailVerifiedCode(uid)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendValidateEmail(email, code)
}

// 注册时验证邮箱
func (domain *VerifiedDomain) VerifiedEmail(uid uint, vcode string) (bool, error) {
	var err error
	var code string
	// 缓存取出
	code, err = domain.cache.GetUserVerifiedEmailCode(uid)
	if err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainCacheGet, DomainName, "GetUserVerifiedEmailCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 生成邮箱验证码
func (domain *VerifiedDomain) genResetPasswordCode(uid uint) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code = XGlobal.RandomString(40)

	// 保存到缓存
	err = domain.cache.SetForgetPasswordVerifiedCode(uid, code)
	if err != nil {
		return "", core.WrapError(err, core.ErrorFormatDomainCacheSet, DomainName, "SetForgetPasswordVerifiedCode")
	}

	return code, nil
}

// 构造验证邮箱邮件
func (domain *VerifiedDomain) sendResetPasswordCodeEmail(email string, code string) error {
	from := "no-reply@" + XGlobal.Config().ServerName
	subject := "Reset Password Code From " + XGlobal.Config().AppName
	// TODO 设置重置密码的链接
	url := XGlobal.Config().ServerName + "?code=" + code
	body := fmt.Sprintf("Click This link To Reset Your Password: %s", url)
	return XMail.SendSimpleMail(from, email, subject, body)
}

// 发送验证码到邮箱
func (domain *VerifiedDomain) SendResetPasswordCodeEmail(dto services.SendEmailForgetPasswordDTO) error {
	var err error
	var code string

	uid := dto.ID
	email := dto.Email

	code, err = domain.genResetPasswordCode(uid)
	if err != nil {
		return err
	}

	// 发送邮件
	return domain.sendResetPasswordCodeEmail(email, code)
}

// 忘记密码时验证重置码
func (domain *VerifiedDomain) VerifiedResetPasswordCode(uid uint, vcode string) (bool, error) {
	var err error
	var code string

	// 缓存取出
	code, err = domain.cache.GetForgetPasswordVerifiedCode(uid)
	if err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainCacheGet, DomainName, "GetForgetPasswordVerifiedCode")
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 生成手机短信验证码
func (domain *VerifiedDomain) genPhoneVerifiedCode(uid uint) (string, error) {
	var err error
	var code string

	// 生成4位随机数字
	code, err = XGlobal.RandomNumber(4)
	if err != nil {
		return "", core.WrapError(err, core.ErrorFormatDomainAlgorithm, DomainName, "global.RandomNumber")
	}

	// 保存到缓存
	err = domain.cache.SetUserVerifiedPhoneCode(uid, code)
	if err != nil {
		return "", core.WrapError(err, core.ErrorFormatDomainCacheSet, DomainName, "SetUserVerifiedPhoneCode")
	}

	return code, nil
}

// 构造短信
func (domain *VerifiedDomain) sendValidatePhoneMsg(phone string, code string) error {

	return nil
}

// 发送验证码到手机短信
func (domain *VerifiedDomain) SendValidatePhoneMsg(dto services.SendPhoneVerifiedCodeDTO) error {
	var err error
	var code string

	uid := dto.ID
	phone := strconv.Itoa(int(dto.Phone))

	code, err = domain.genPhoneVerifiedCode(uid)
	if err != nil {
		return err
	}

	return domain.sendValidatePhoneMsg(phone, code)
}

// 注册时验证手机短信
func (domain *VerifiedDomain) VerifiedPhone(uid uint, vcode string) (bool, error) {
	var err error
	var code string

	// 缓存取出
	code, err = domain.cache.GetUserVerifiedPhoneCode(uid)
	if err != nil {
		return false, err
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}
