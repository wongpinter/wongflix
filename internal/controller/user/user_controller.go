package userRepo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wongpinter/wongflix/internal/entity/user"
)

type UserAuthController interface {
	PostSignUp(response http.ResponseWriter, request *http.Request)
	PostLogin(response http.ResponseWriter, request *http.Request)
}

var (
	service user.UseCase
)

type controller struct{}

func NewUserController(useCase user.UseCase) UserAuthController {
	service = useCase
	return &controller{}
}

func (*controller) PostSignUp(response http.ResponseWriter, request *http.Request) {
	u := user.Register{}

	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(response, "Error decoding response object", http.StatusBadRequest)
		return
	}

	_, err := service.SignUp(context.Background(), &u)

	user, err := json.Marshal(&u)

	if err != nil {
		fmt.Println(err)
		http.Error(response, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(user)
}

func (*controller) PostLogin(response http.ResponseWriter, request *http.Request) {
	var u struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(response, "Error decoding response object", http.StatusBadRequest)
		return
	}

	token, err := service.Login(context.Background(), u.Email, u.Password)

	fmt.Println(token)

	user, err := json.Marshal(&u)

	if err != nil {
		fmt.Println(err)
		http.Error(response, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(user)
}
