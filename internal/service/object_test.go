package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/object/v1"
	"testing"
)

var objectClient v1.ObjectClient

func init() {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		panic(err)
	}
	objectClient = v1.NewObjectClient(conn)
}

func TestOrderService_CreateOrder(t *testing.T) {
	object, err := objectClient.CreateObject(ctx, &v1.CreateObjectRequest{
		Name:     "test1",
		Size:     1024,
		BucketId: 1463584877990383616,
		Type:     "application/json",
		Replicas: 3,
	})
	if err != nil {
		t.Error(err)
	}
	if object.ObjectId == 0 {
		t.Error("create object failed")
	}
	t.Logf("%+v", object)
}

func TestOrderService_GetObject(t *testing.T) {
	resp, err := objectClient.GetObject(ctx, &v1.GetObjectRequest{
		ObjectId: 8769702054114365440,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}

func TestOrderService_ListObject(t *testing.T) {
	resp, err := objectClient.ListObject(ctx, &v1.ListObjectRequest{
		BucketId: 1463584877990383616,
		Page:     0,
		PageSize: 10,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}

func TestOrderService_DeleteObject(t *testing.T) {
	resp, err := objectClient.DeleteObject(ctx, &v1.DeleteObjectRequest{
		ObjectId: 8769702054114365440,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
