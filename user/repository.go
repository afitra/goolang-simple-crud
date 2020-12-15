package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByUserName(username string) (User, error)
	FindByID(ID int) (User, error)
	UpdateUser(user User) (User, error)
	GetAllUser() ([]User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByUserName(username string) (User, error) {

	var user User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindByID(ID int) (User, error) {

	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {

		return user, err
	}
	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetAllUser() ([]User, error) {

	var allUser []User
	err := r.db.Find(&allUser).Error
	if err != nil {
		return allUser, err
	}
	return allUser, nil
}

func (r *repository) Update(user User) (User, error) {

	err := r.db.Save(&user).Error

	if err != nil {

		return user, err
	}
	return user, nil

}
