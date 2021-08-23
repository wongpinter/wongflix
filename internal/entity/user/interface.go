package user

import "context"

type Authentication interface {
	Login(ctx context.Context, email string, password string) (token string, err error)
	SignUp(ctx context.Context, register *Register) (*Public, error)
	// LogOut(ctx context.Context, id UUID) error
}

type Read interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type Write interface {
	Store(ctx context.Context, user *User) (*User, error)
}

type ModelDB interface {
	ToEntity() (*User, error)
	FromEntity(user *User) (*User, error)
}

type UseCase interface {
	Authentication
	Read
}

type Repository interface {
	Read
	Write
}
