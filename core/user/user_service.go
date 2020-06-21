package user

import (
	"GoWebScaffold/services"
	"sync"
)

// 服务层，实现services包定义的服务，在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.IUserService = new(userService)
var once sync.Once

func init() {
	// 初始化该业务模块时实例化服务
	once.Do(func() {
		services.SetUserService(new(userService))
	})
}

type userService struct{}

func (*userService) CreateUser(dto services.CreateUserDTO) (*services.UserDTO, error) {
	panic("implement me")
}

func (*userService) GetUserInfo(userId string) (*services.UserDTO, error) {
	panic("implement me")
}

func (*userService) SetUserInfo(dto services.SetUserInfoDTO) error {
	panic("implement me")
}
