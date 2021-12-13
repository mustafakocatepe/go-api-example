package user

type createUserRequest struct {
	Username string `json:"username"`
	Surname  string `json:"surname"`
}

type updateUserNameRequest struct {
	Username string `json:"username"`
}

type updateUserRequest struct {
	Username string `json:"username"`
	Surname  string `json:"surname"`
}
