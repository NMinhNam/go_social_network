package service

import "go-social-network/internal/repository"

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *UserService) GetUserInfo() string {
	return us.userRepository.GetUserInfo()
}
