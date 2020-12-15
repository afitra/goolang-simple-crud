package user

import "golang.org/x/crypto/bcrypt"

type service struct {
	repository Repository // repo yg hutruf kecil itu variable punya service sendiri bukan punya Repository
}

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}
	user.Username = input.Username

	user.Nama_Lengkap = input.Nama_Lengkap
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
