package user

const (
	DomainName = "UserDomain"
	// 用户状态相关
	UserStatusNotVerified  = iota // 未验证 0
	UserStatusNormal              // 已验证 1
	UserStatusDeactivation        // 已停用 2

)
