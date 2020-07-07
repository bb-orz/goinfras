package mail

// 用户邮箱验证
const (
	UserCacheVerifiedEmailCodePrefix = "user.verified.email.code." // 缓存邮箱验证码key前缀
	UserCacheVerifiedEmailCodeExpire = 60 * 60 * 3                 // 缓存邮箱验证码超时时间
)
