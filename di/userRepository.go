package di

type UserRepository interface {
	All() ([]User, error)
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct{}
