package user

type userService struct {
	userRepository userRepositoryInterface
}

func NewUserService(r userRepositoryInterface) userServiceInterface {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Register(u User) error {
	return s.userRepository.Create(u)
}
