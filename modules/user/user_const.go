package user

const (
	// 用户状态相关
	UserStatusNotVerified  = iota // 未验证 0
	UserStatusNormal              // 已验证 1
	UserStatusDeactivation        // 已停用 2

	// 用户邮箱手机验证
	UserCacheVerifiedEmailCodePrefix = "user.verified.email.code." // 缓存邮箱验证码key前缀
	UserCacheVerifiedEmailCodeExpire = 60 * 60 * 3                 // 缓存邮箱验证码超时时间
	UserCacheVerifiedPhoneCodePrefix = "user.verified.phone.code." // 缓存手机验证码key前缀
	UserCacheVerifiedPhoneCodeExpire = 60 * 5                      // 缓存手机验证码超时时间

)
