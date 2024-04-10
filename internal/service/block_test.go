package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/block/v1"
	"testing"
)

var ctx = context.Background()
var blockClient v1.BlockClient

func init() {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		panic(err)
	}
	blockClient = v1.NewBlockClient(conn)
}

func TestCreateBlock(t *testing.T) {
	resp, err := blockClient.CreateBlock(ctx, &v1.CreateBlockRequest{
		ObjectId: 8769702054114365440,
		Size:     10224,
		Num:      1,
		Checksum: 131323,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("resp: %v", resp)
}

func TestListBlock(t *testing.T) {
	resp, err := blockClient.ListBlock(ctx, &v1.ListBlockRequest{
		ObjectId: 8769702054114365440,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("resp: %v", resp)
}

func TestDeleteBlock(t *testing.T) {
	resp, err := blockClient.DeleteBlock(ctx, &v1.DeleteBlockRequest{
		ObjectId: 8769702054114365440,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("resp: %v", resp)
}
