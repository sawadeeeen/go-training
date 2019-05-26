package di

type UserService struct {
	repo UserRepository
}

func (s *UserService) All() ([]User, error) {
	return s.repo.All()
}
