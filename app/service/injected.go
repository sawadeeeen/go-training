package service

type Injected struct{}

func (Injected) call(arg string) string {
	return "injected!"
}
