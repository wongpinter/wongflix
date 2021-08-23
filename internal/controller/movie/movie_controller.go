package movie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/wongpinter/wongflix/internal/entity/movie"
)

type MovieController interface {
	Search(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
	GetByTitle(response http.ResponseWriter, request *http.Request)
}

var (
	service movie.UseCase
)

type controller struct{}

func NewMovieController(useCase movie.UseCase) MovieController {
	service = useCase
	return &controller{}
}

func (*controller) Search(response http.ResponseWriter, request *http.Request) {
	query := request.FormValue("query")
	page := request.FormValue("page")

	i, _ := strconv.Atoi(page)

	res, _ := service.Search(context.Background(), query, i)

	search, err := json.Marshal(res)

	if err != nil {
		return
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(search)
}

func (*controller) GetByID(response http.ResponseWriter, request *http.Request) {
	movieID := chi.URLParam(request, "id")

	res, _ := service.GetByID(context.Background(), movieID)

	movie, err := json.Marshal(res)

	if err != nil {
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(movie)
}

func (*controller) GetByTitle(response http.ResponseWriter, request *http.Request) {
	movieTitle := chi.URLParam(request, "title")

	res, _ := service.GetByTitle(context.Background(), movieTitle)

	movie, err := json.Marshal(res)

	if err != nil {
		fmt.Println(err)
		http.Error(response, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(movie)
}
