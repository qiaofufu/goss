package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/bucket/v1"
	"testing"
)

func TestBucketService_CreateBucket(t *testing.T) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		t.Fatal(err)
	}
	client := v1.NewBucketClient(conn)

	bucket, err := client.CreateBucket(context.Background(), &v1.CreateBucketRequest{
		Name:        "test-bucket",
		Description: "test-description",
		OwnerID:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("bucket: %+v", bucket)
}
