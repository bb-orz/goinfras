package mail

import (
	"GoWebScaffold/infras/global"
	"GoWebScaffold/infras/mail"
	"GoWebScaffold/services"
	"fmt"
)

/*领域层：实现具体业务逻辑*/
type MailDomain struct {
	cache *mailCache
}

func NewMailDomain() *MailDomain {
	domain := new(MailDomain)
	domain.cache = NewMailCache()
	return domain
}

// 生成验证码
func (domain *MailDomain) genEmailVerifiedCode(uid int) (string, error) {
	// 生成6位随机字符串
	code := global.RandomString(6)

	// 保存到缓存
	err := domain.cache.SetUserVerifiedEmailCode(uid, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (domain *MailDomain) sendEmail(email string, code string) error {
	from := "no-reply@" + global.Config().ServerName
	subject := "Verified Email Code From " + global.Config().AppName
	body := fmt.Sprintf("Verified Code: %s", code)
	return mail.SendSimpleMail(from, email, subject, body)
}

// 发送验证码到邮箱
func (domain *MailDomain) SendEmail(dto services.SendEmailForVerifiedDTO) error {
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
func (domain *MailDomain) VerifiedEmail(uid int, vcode string) (bool, error) {

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
