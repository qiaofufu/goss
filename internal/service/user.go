package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"master/api/user/v1"
	"master/internal/biz"
	"time"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func (s *UserService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginReply, error) {
	token, err := s.uc.Login(ctx, request.Username, request.Password)
	if err != nil {
		s.log.Errorf("Service Login Error: %v", err)
		return &user.LoginReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &user.LoginReply{
		Status:    0,
		Msg:       "Success",
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}, nil
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	user := &biz.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Roles:    req.Roles,
	}
	err := s.uc.CreateUser(ctx, user, req.OperatorId)
	if err != nil {
		s.log.Errorf("Service Create User Error: %v", err)
		return &v1.CreateUserReply{
			Id:     0,
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &v1.CreateUserReply{
		Id:     user.Id,
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	user := &biz.User{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Active:   req.Active,
		Roles:    req.Roles,
	}
	err := s.uc.UpdateUser(ctx, user, req.OperatorId)
	if err != nil {
		s.log.Errorf("Service Update User Error: %v", err)
		return &v1.UpdateUserReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &v1.UpdateUserReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := s.uc.DeleteUser(ctx, req.Id, req.OperatorId)
	if err != nil {
		s.log.Errorf("Service Delete User Error: %v", err)
		return &v1.DeleteUserReply{
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	return &v1.DeleteUserReply{
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.GetUserByID(ctx, req.Id, req.OperatorId)
	if err != nil {
		s.log.Errorf("Service Get User Error: %v", err)
		return &v1.GetUserReply{
			User:   nil,
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}

	return &v1.GetUserReply{
		User: &v1.UserInfo{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Roles,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Status: 0,
		Msg:    "Success",
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	users, total, err := s.uc.ListUser(ctx, int(req.Page), int(req.Size), req.OperatorId)
	if err != nil {
		s.log.Errorf("Service List User Error: %v", err)
		return &v1.ListUserReply{
			Users:  nil,
			Total:  0,
			Status: 1,
			Msg:    err.Error(),
		}, nil
	}
	var pbUsers []*v1.UserInfo
	for _, user := range users {
		pbUsers = append(pbUsers, &v1.UserInfo{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Roles,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return &v1.ListUserReply{
		Users:  pbUsers,
		Total:  total,
		Status: 0,
		Msg:    "Success",
	}, nil
}
