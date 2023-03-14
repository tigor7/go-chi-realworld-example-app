package user

type profileResponse struct {
	Profile profile `json:"profile"`
}

type profile struct {
	Username string  `json:"username"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
}

func NewProfileResponse(u User) profileResponse {
	p := profile{
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
	}
	return profileResponse{
		Profile: p,
	}
}
