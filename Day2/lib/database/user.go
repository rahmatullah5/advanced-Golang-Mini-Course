package database

import (
	"Day2/config"
	"Day2/models"
)

func GetUsers() (interface{}, error) {

	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUser(id string, userValues models.User) (interface{}, error) {

	config.DB.Create(&userValues)

	return userValues, nil
}

func GetUser(id string) (interface{}, error) {
	var user models.User

	e := config.DB.First(&user, id).Error

	if e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUser(id string, userValues models.User) (interface{}, error) {
	var user models.User

	e := config.DB.First(&user, id).Error

	if e != nil {
		return nil, e
	}

	config.DB.Model(&user).Updates(userValues)
	return user, nil
}

func DeleteUser(id string) (interface{}, error) {
	var user models.User

	e := config.DB.First(&user, id).Error

	if e != nil {
		return nil, e
	}

	config.DB.Delete(&user, id)

	return user, nil
}
