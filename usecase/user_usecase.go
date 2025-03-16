package usecase

import (
	"github.comElcioAngelo/TrabalhoTCC.git/model"
	"github.comElcioAngelo/TrabalhoTCC.git/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserRepository(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
} 

func (su *UserUseCase) GetUser(user_id string) (model.User,error) {
	return su.repository.GetUser(user_id)
}

func (su *UserUseCase) CreateUser(user model.User) (error) {
	return su.repository.CreateUser(user)
}






