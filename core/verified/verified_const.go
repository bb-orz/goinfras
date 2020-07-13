package verified

// 用户邮箱验证
const (
	DomainName                       = "VerifiedDomain"
	UserCacheVerifiedEmailCodePrefix = "user.verified.email.code." // 缓存邮箱验证码key前缀
	UserCacheVerifiedEmailCodeExpire = 60 * 60 * 3                 // 缓存邮箱验证码超时时间

	UserCacheVerifiedPhoneCodePrefix = "user.verified.phone.code." // 缓存手机验证码key前缀
	UserCacheVerifiedPhoneCodeExpire = 60 * 5                      // 缓存手机验证码超时时间
)
