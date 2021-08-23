package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/wongpinter/wongflix/internal/entity/user"
	auth "github.com/wongpinter/wongflix/pkg/jwt"
	"github.com/wongpinter/wongflix/pkg/password"
)

type userUseCase struct {
	ur user.Repository
}

var (
	pwd = password.NewHash()
	jwt = auth.NewJWT()
)

func NewUserUC(userRepo user.Repository) user.UseCase {
	return &userUseCase{
		ur: userRepo,
	}
}

func (userUC *userUseCase) GetByID(ctx context.Context, id string) (*user.User, error) {
	if id == "" {
		return nil, errors.New("user id is required")
	}

	return userUC.ur.GetByID(ctx, id)
}

func (userUC *userUseCase) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	if email == "" {
		return nil, errors.New("user email is required")
	}

	return userUC.ur.GetByEmail(ctx, email)
}

func (userUC *userUseCase) Login(ctx context.Context, email string, password string) (token string, err error) {
	details, err := userUC.ur.GetByEmail(ctx, email)
	if err != nil {
		return token, err
	}

	if details.Email == "" {
		return token, errors.New(fmt.Sprintf("user %s does not exists", email))
	}

	err = pwd.Compare(details.Password, password)

	if err != nil {
		return token, errors.New(fmt.Sprintf("user %s password incorrect", email))
	}

	token, err = jwt.Generate(&user.JWT{
		ID:    details.ID,
		Email: details.Email,
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (userUC *userUseCase) SignUp(ctx context.Context, register *user.Register) (*user.Public, error) {

	hashed, err := pwd.Generate(register.Password)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to encrypt password with error: %s", err))
	}

	details, _ := userUC.ur.GetByEmail(ctx, register.Email)
	// if err != nil {
	// 	return nil, errors.New(fmt.Sprintf("failed to get user with error: %s", err))
	// }

	if details.Email == register.Email {
		return nil, errors.New(fmt.Sprintf("user already exists"))
	}

	result, err := userUC.ur.Store(ctx, &user.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: hashed,
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to add user with errors: %s", err))
	}

	return &user.Public{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}, nil
}
