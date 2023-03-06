package user

type userService struct {
	userRepository userRepository
}

func NewUserService(r userRepository) userService {
	return userService{
		userRepository: r,
	}
}
