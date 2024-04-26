package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"master/third_party/jwt"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Email     string
	CreatedAt int64
	UpdatedAt int64
	Active    bool
	Roles     []string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id int64) (*User, error)
	ListUser(ctx context.Context, page int, size int) ([]*User, int64, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	VerifyUserLogin(ctx context.Context, username, password string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *User, operatorID int64) error {
	// TODO: check operatorID permission

	err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUsecase.CreateUser: %w", err)
	}
	return nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User, operatorID int64) error {
	// TODO: check operatorID permission

	err := uc.repo.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUsecase.UpdateUser: %w", err)
	}
	return nil
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int64, operatorID int64) (*User, error) {
	// TODO: check operatorID permission

	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserUsecase.GetUserByID: %w", err)
	}
	return user, nil
}

func (uc *UserUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := uc.repo.VerifyUserLogin(ctx, username, password)
	if err != nil {
		return "", fmt.Errorf("UserUsecase.Login: %w", err)
	}
	// TODO: generate token
	token, err := jwt.GenToken(user.Id)
	if err != nil {
		return "", fmt.Errorf("UserUsecase.Login: %w", err)
	}
	return token, nil
}

func (uc *UserUsecase) ListUser(ctx context.Context, page, size int, operatorID int64) ([]*User, int64, error) {
	// TODO: check operatorID permission

	user, total, err := uc.repo.ListUser(ctx, page, size)
	if err != nil {
		return nil, 0, fmt.Errorf("UserUsecase.ListUser: %w", err)
	}
	return user, total, nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64, operatorID int64) error {
	// TODO: check operatorID permission

	err := uc.repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("UserUsecase.DeleteUser: %w", err)
	}
	return nil
}
