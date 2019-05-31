package main

type UserRepository struct {
	database *sql.DB
}

func (repository *UserRepository) FindById(uid int) *User {

}

func (repository *UserRepository) FindAll() []*User {
	rows, _ := repository.database.Query()
}
