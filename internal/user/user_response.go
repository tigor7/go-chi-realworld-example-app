package user

type userResponse struct {
	User user `json:"user"`
}

type user struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func NewUserResponse(u User, token string) userResponse {
	us := user{
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    token,
	}
	return userResponse{
		User: us,
	}
}
