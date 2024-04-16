package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/user/v1"
	"testing"
)

var (
	client v1.UserClient
)

func init() {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		panic(err)
	}
	client = v1.NewUserClient(conn)
}

func TestUserService_CreateUser(t *testing.T) {
	user, err := client.CreateUser(ctx, &v1.CreateUserRequest{
		Username:   "qiaohh",
		Password:   "123456",
		Email:      "707402933@qq.com",
		Roles:      []string{"admin"},
		OperatorId: 1000,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestUserService_GetUser(t *testing.T) {
	user, err := client.GetUser(ctx, &v1.GetUserRequest{
		Id: 906783175634587648,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestUserService_GetUsers(t *testing.T) {
	users, err := client.ListUser(ctx, &v1.ListUserRequest{
		Size: 10,
		Page: 1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(users)
}

func TestUserService_UpdateUser(t *testing.T) {
	resp, err := client.UpdateUser(ctx, &v1.UpdateUserRequest{
		Id:         906783175634587648,
		Username:   "test-2",
		Password:   "8023qiao",
		OperatorId: 1000,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUserService_DeleteUser(t *testing.T) {
	resp, err := client.DeleteUser(ctx, &v1.DeleteUserRequest{
		Id: 1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
