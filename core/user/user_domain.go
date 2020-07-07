package user

import (
	"GoWebScaffold/infras/global"
	"GoWebScaffold/infras/mail"
	"GoWebScaffold/services"
	"fmt"
	"github.com/segmentio/ksuid"
	"strconv"
)

/*领域层：实现具体业务逻辑*/
type UserDomain struct {
	dao   *UserDAO
	cache *userCache
}

func NewUserDomain() *UserDomain {
	domain := new(UserDomain)
	domain.dao = NewUserDao()
	domain.cache = NewUserCache()
	return domain
}

// 生成用户编号
func (domain *UserDomain) generateUserNo() string {
	// 采用ksuid的ID生成策略来创建No
	// 全局唯一的ID
	return ksuid.New().Next().String()
}

// 加密密码，设置密文和盐值到po
func (domain *UserDomain) encryptPassword(password string) (hashStr, salt string) {
	hashStr, salt = global.HashPassword(password)
	return
}

// 查找用户是否已存在
func (domain *UserDomain) IsUserEmailExist(dto services.CreateUserWithEmailDTO) (bool, error) {
	if isExist, err := domain.dao.IsEmailExist(dto.Email); err != nil {
		fmt.Println("Error:", err)
		return false, err
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 邮箱账号创建用户
func (domain *UserDomain) CreateUserForEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	userModel := User{}
	userModel.Name = dto.Name
	userModel.Email = dto.Email
	userModel.No = domain.generateUserNo()
	userModel.Password, userModel.Salt = domain.encryptPassword(dto.Password)
	userModel.Status = UserStatusNotVerified // 初始创建时未验证状态
	fmt.Println("UserModel:", userModel)
	var user *User
	var err error
	if user, err = domain.dao.Create(userModel); err != nil {
		return nil, err
	}

	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfo(uid int) (*services.UserDTO, error) {
	user, err := domain.dao.GetById(uid)
	if err != nil {
		return nil, err
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid int, field string, value interface{}) error {
	return domain.dao.SetUserInfo(uid, field, value)
}

// 设置多个用户信息
func (domain *UserDomain) SetUserInfos(uid int, dto services.SetUserInfoDTO) error {
	return domain.dao.SetUserInfos(uid, dto)
}

// 生成验证码
func (domain *UserDomain) genEmailVerifiedCode(uid int) (string, error) {
	// 生成6位随机字符串
	code := global.RandomString(6)

	// 保存到缓存
	err := domain.cache.SetUserVerifiedEmailCode(uid, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (domain *UserDomain) sendEmail(email string, code string) error {
	from := "no-reply@" + global.Config().ServerName
	subject := "Verified Email Code From " + global.Config().AppName
	body := fmt.Sprintf("Verified Code: %s", code)
	return mail.SendSimpleMail(from, email, subject, body)
}

// 发送验证码到邮箱
func (domain *UserDomain) SendEmail(dto services.SendEmailForVerifiedDTO) error {
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
func (domain *UserDomain) VerifiedEmail(uid int, vcode string) (bool, error) {

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

// 生成验证码
func (domain *UserDomain) genPhoneVerifiedCode(uid int) (string, error) {
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

func (domain *UserDomain) sendPhoneMsg(phone string, code string) error {

	return nil
}

// 发送验证码到手机短信
func (domain *UserDomain) SendPhoneMsg(dto services.SendPhoneVerifiedCodeDTO) error {
	uid := int(dto.ID)
	phone := strconv.Itoa(int(dto.Phone))

	code, err := domain.genPhoneVerifiedCode(uid)
	if err != nil {
		return err
	}

	return domain.sendPhoneMsg(phone, code)
}

// 验证邮箱
func (domain *UserDomain) VerifiedPhone(uid int, vcode string) (bool, error) {
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
