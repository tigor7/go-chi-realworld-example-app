package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tigor7/go-chi-realworld-example-app/internal/auth"
)

type userService struct {
	userRepository userRepositoryInterface
}

func NewUserService(r userRepositoryInterface) userServiceInterface {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Register(u User) (string, error) {
	u.ID = uuid.New()
	err := s.userRepository.Create(u)
	if err != nil {
		return "", err
	}

	token, err := auth.CreateJWT(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) Login(u User) (User, string, error) {
	us, err := s.userRepository.GetByEmail(u.Email)
	if err != nil {
		return us, "", err
	}
	if !ComparePassword(us.Password, u.Password) {
		return us, "", errors.New("Username and password do not match")
	}
	token, err := auth.CreateJWT(us.ID)
	if err != nil {
		return us, "", err
	}
	return us, token, nil
}

func (s *userService) GetProfile(username string) (User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s *userService) GetUserByID(id uuid.UUID) (User, error) {
	return s.userRepository.GetByUserID(id)
}

func (s *userService) Follow(uid uuid.UUID, username string) (User, error) {
	friend, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return friend, err
	}
	return friend, s.userRepository.Follow(uid, friend.ID)
}

func (s *userService) Unfollow(uid uuid.UUID, username string) (User, error) {
	friend, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return friend, err
	}
	return friend, s.userRepository.Unfollow(uid, friend.ID)
}
