package user

type createUserRequest struct {
	Username string `json:"username"`
	Surname  string `json:"surname"`
}

type updateUserRequest struct {
	Username string `json:"username"`
}
