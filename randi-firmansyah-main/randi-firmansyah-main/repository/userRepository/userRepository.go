package userRepository

import (
	"randi_firmansyah/connections/database"
	"randi_firmansyah/models/userModel"
)

var DB = database.Connected

func FindAll() ([]userModel.User, error) {
	var users []userModel.User
	err := DB.Find(&users).Error
	return users, err
}

func FindByID(ID int) (userModel.User, error) {
	var user userModel.User
	err := DB.First(&user, ID).Error
	return user, err
}

func Create(user userModel.User) (userModel.User, error) {
	err := DB.Create(&user).Error
	return user, err
}

func UpdateV2(user userModel.User) (userModel.User, error) {
	err := DB.Save(&user).Error
	return user, err
}

func Update(id int, user userModel.User) (userModel.User, error) {
	err := DB.Model(userModel.User{}).Where("id = ?", id).Updates(user).Error
	return user, err
}

func Delete(user userModel.User) (userModel.User, error) {
	err := DB.Delete(&user).Error
	return user, err
}
