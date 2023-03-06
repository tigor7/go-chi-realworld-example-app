package user

type userService struct {
	userRepository userRepositoryInterface
}

func NewUserService(r userRepositoryInterface) userServiceInterface {
	return &userService{
		userRepository: r,
	}
}
