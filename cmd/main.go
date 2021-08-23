package main

import (
	"fmt"
	"net/http"

	"gorm.io/gorm/logger"

	"github.com/wongpinter/wongflix/config"

	"github.com/wongpinter/wongflix/internal/infrastructure/datasource"
	"github.com/wongpinter/wongflix/internal/infrastructure/router"

	userCtrl "github.com/wongpinter/wongflix/internal/controller/user"
	userRepo "github.com/wongpinter/wongflix/internal/repository/user"
	userUC "github.com/wongpinter/wongflix/internal/usecase/user"

	movieCtrl "github.com/wongpinter/wongflix/internal/controller/movie"
	movieRepo "github.com/wongpinter/wongflix/internal/repository/movie"
	movieUC "github.com/wongpinter/wongflix/internal/usecase/movie"
)

var (
	db, _ = datasource.Connect(
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_DATABASE,
		logger.Error,
	)

	userService = userUC.NewUserUC(
		userRepo.NewMysqlRepository(db),
	)

	userController = userCtrl.NewUserController(userService)

	movieService = movieUC.NewMovieUseCase(
		movieRepo.NewOmdbAPIRepository(),
	)

	movieController = movieCtrl.NewMovieController(movieService)

	httpRouter = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Up and Running")
	})

	httpRouter.POST("/sign-up", userController.PostSignUp)
	httpRouter.POST("/sign-in", userController.PostLogin)

	httpRouter.GET("/v1/movie/search", movieController.Search)
	httpRouter.GET("/v1/movie/id/{id}", movieController.GetByID)
	httpRouter.GET("/v1/movie/title/{title}", movieController.GetByTitle)

	httpRouter.SERVE(config.API_PORT)
}
