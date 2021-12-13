package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mustafakocatepe/go-api-example/handler/api/errors"
	"github.com/mustafakocatepe/go-api-example/handler/render"
	"github.com/mustafakocatepe/go-api-example/model"
	userService "github.com/mustafakocatepe/go-api-example/service/user"
)

func HandleUsers(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, result := userService.GetUsers(model.UserArray)
		if !result {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, users, http.StatusOK)
	}
}

func HandleUpdate(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")

		if len(id) == 0 {
			render.BadRequest(w, errors.New(""))
			return
		}

		var req updateUserRequest
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

		var user model.User

		if len(req.Username) != 0 {
			user.UserName = req.Username
		}

		if len(req.Surname) != 0 {
			user.Surname = req.Surname
		}

		result := userService.UpdateUserByUserId(model.UserArray, id, user)

		if !result {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, "", http.StatusNoContent)
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

		user, found := userService.GetUserByUserName(users, userName)
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

		render.JSON(w, user, http.StatusCreated)
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

		result := userService.DeleteUserByUserId(users, id)

		if !result {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, "", http.StatusNoContent)

	}
}

func HandleUpdateUserName(us model.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")

		if len(id) == 0 {
			render.BadRequest(w, errors.New(""))
			return
		}

		var req updateUserNameRequest
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

		users := model.UserArray

		result := userService.UpdateUserNameByUserId(users, id, req.Username)

		if !result {
			render.NotFound(w, errors.New("Kullanici bulunamadi"))
			return
		}

		render.JSON(w, "", http.StatusNoContent)
	}
}
