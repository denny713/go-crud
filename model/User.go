package model

import "Github/go-crud/config"

func GetAllUser(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetByUsername(user *User, id string) (err error) {
	if err = config.DB.Where("username = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	if err = config.DB.Where("username = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
