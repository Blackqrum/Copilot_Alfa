package modules

type User struct {
	ID       int
	Email    string
	Password string
	Name     string
	CreateAt string
}

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
}

type LoginRequest struct {
	Email    string
	Password string
}
