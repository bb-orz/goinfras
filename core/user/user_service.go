package user

import "GoWebScaffold/services"

// 服务层，实现services包定义的服务，在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑
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
