package models

type User struct {
	ID int64 `json:"id"`

	Email string `json:"email"`
	Login string `json:"login"`
	Name  string `json:"name"`

	Salt []byte `json:"-"`
	Hash []byte `json:"-"`
}

type UserRequest struct {
	User
	ID       int64  `json:"-"`
	Password string `json:"password"`
}

type UserResponse struct {
	User
}
