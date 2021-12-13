package user

import (
	"strconv"

	"github.com/mustafakocatepe/go-api-example/model"
)

func GetUserByUserName(slice []model.User, val string) (*model.User, bool) {
	for _, item := range slice {
		if item.UserName == val && item.IsActive {
			return &item, true
		}
	}
	return nil, false
}
func GetUsers(slice []model.User) ([]model.User, bool) {
	var responseModel []model.User
	for _, item := range slice {
		if item.IsActive {
			responseModel = append(responseModel, item)
		}
	}
	return responseModel, true
}
func DeleteUserByUserId(slice []model.User, val string) bool {
	for i, item := range slice {
		value, _ := strconv.Atoi(val)
		if item.UserId == value && item.IsActive {
			(model.UserArray[i]).IsActive = false
			return true
		}
	}
	return false
}
func UpdateUserNameByUserId(slice []model.User, val string, userName string) bool {
	for i, item := range slice {
		value, _ := strconv.Atoi(val)
		if item.UserId == value && item.IsActive {
			(model.UserArray[i]).UserName = userName
			return true
		}
	}
	return false
}
