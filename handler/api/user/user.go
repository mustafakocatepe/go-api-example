package user

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mustafakocatepe/go-api-example/handler/api/errors"
	"github.com/mustafakocatepe/go-api-example/handler/render"
	"github.com/mustafakocatepe/go-api-example/model"
)

func HandleUsers(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, result := GetUsers(model.UserArray)
		if !result {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, users, http.StatusOK)
	}
}

func HandleUpdate(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := model.User{UserId: 2, UserName: "msk", Surname: "Kocatepe"}

		render.JSON(w, user, http.StatusOK)
	}
}

func HandleUser(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userName := chi.URLParam(r, "username")

		if len(userName) == 0 {
			render.BadRequest(w, errors.New(""))
			return
		}

		users := model.UserArray

		user, found := GetUserByUserName(users, userName)
		if !found {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, user, http.StatusOK)

	}
}

func HandleCreateUser(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req createUserRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &req); err != nil {
			render.BadRequest(w, err)
			return
		}

		user := model.User{
			UserId:   len(model.UserArray) + 1,
			UserName: req.Username,
			Surname:  req.Surname,
			IsActive: true,
		}

		model.UserArray = append(model.UserArray, user)

		render.JSON(w, model.UserArray, http.StatusCreated)
	}
}

func HandleDelete(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")

		if len(id) == 0 {
			render.BadRequest(w, errors.New(""))
			return
		}

		users := model.UserArray

		result := DeleteUserByUserId(users, id)

		if !result {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, model.UserArray, http.StatusOK) //TODO: DEGISTIR

	}
}

//-----------------------------------------------------
func GetUserByUserName(slice []model.User, val string) (*model.User, bool) {
	for _, item := range slice {
		if item.UserName == val {
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
