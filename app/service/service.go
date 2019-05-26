package service

type Service interface {
	Call(arg string) string
}

type ServiceImpl struct {
	Injected *Injected
}

func (s *ServiceImpl) Call(arg string) string {
	return "mainService" + arg
}
