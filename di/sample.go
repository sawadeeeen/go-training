package di

container := NewContainer()

rd := &Definition{
	Name: "UserRepository",
	Builder: func(c *Container) interface{} {
		return NewUserRepository()
	},
}
container.Register("UserRepository", rd)

sd := &Definition{
	Name := "UserService",
	Builder: func(c *Container) interface{} {
		repo, _ := c.Get("UserRepository").(UserRepository)
		return NewUserService(repo)
	}, 
}
container.Register("UserService", sd)

service := container.Get("UserRepository")
