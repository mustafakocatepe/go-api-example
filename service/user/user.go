package user

import (
	"strconv"

	"github.com/mustafakocatepe/go-api-example/model"
)

func GetUserByUserName(users []model.User, id string) (*model.User, bool) {
	for _, item := range users {
		if item.UserName == id && item.IsActive {
			return &item, true
		}
	}
	return nil, false
}
func GetUsers(users []model.User) ([]model.User, bool) {
	var responseModel []model.User
	for _, item := range users {
		if item.IsActive {
			responseModel = append(responseModel, item)
		}
	}
	return responseModel, true
}
func DeleteUserByUserId(users []model.User, id string) bool {
	for i, item := range users {
		value, _ := strconv.Atoi(id)
		if item.UserId == value && item.IsActive {
			(model.UserArray[i]).IsActive = false
			return true
		}
	}
	return false
}
func UpdateUserNameByUserId(users []model.User, id string, userName string) bool {
	for i, item := range users {
		value, _ := strconv.Atoi(id)
		if item.UserId == value && item.IsActive {
			(model.UserArray[i]).UserName = userName
			return true
		}
	}
	return false
}

func UpdateUserByUserId(users []model.User, id string, user model.User) bool {
	for i, item := range users {
		value, _ := strconv.Atoi(id)
		if item.UserId == value && item.IsActive {

			if len(user.UserName) != 0 {
				(model.UserArray[i]).UserName = user.UserName
			}

			if len(user.Surname) != 0 {
				(model.UserArray[i]).Surname = user.Surname
			}
			return true
		}
	}
	return false
}
