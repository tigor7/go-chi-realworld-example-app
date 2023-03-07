package user

import "github.com/google/uuid"

type userService struct {
	userRepository userRepositoryInterface
}

func NewUserService(r userRepositoryInterface) userServiceInterface {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Register(u User) error {
	u.ID = uuid.New()
	return s.userRepository.Create(u)
}
