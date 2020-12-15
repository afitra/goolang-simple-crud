package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
}

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(input int) (User, error)
	GetAllUser() ([]User, error)
	SaveProfile(id int, fileLocation string) (User, error)
	UpdateUser(inputID GetUserDetailByID, inputData RegisterUserInput, sessionID int) (User, error)
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

func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUserName(username)
	if err != nil {

		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User tidak ada")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) GetUserByID(input int) (User, error) {
	user, err := s.repository.FindByID(input)

	if err != nil {

		return user, err
	}
	return user, nil
}

func (s *service) GetAllUser() ([]User, error) {
	user, err := s.repository.GetAllUser()

	if err != nil {

		return user, err
	}
	return user, nil
}

func (s *service) SaveProfile(id int, fileLocation string) (User, error) {

	user, err := s.repository.FindByID(id)
	if err != nil {

		return user, err
	}

	user.Foto = fileLocation
	updateUser, err := s.repository.UpdateUser(user)
	if err != nil {

		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) UpdateUser(inputID GetUserDetailByID, inputData RegisterUserInput, sessionID int) (User, error) {

	user, err := s.repository.FindByID(inputID.ID)
	if err != nil {

		return user, err

	}

	if inputID.ID != sessionID {
		fmt.Println("iniiiii dia")
		return user, errors.New("you bukan pemilik user")

	}

	user.Username = inputData.Username

	user.Nama_Lengkap = inputData.Nama_Lengkap

	hash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)

	updateUser, err := s.repository.Update(user)

	if err != nil {
		return updateUser, err
	}
	return updateUser, nil

}
