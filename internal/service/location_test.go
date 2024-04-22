package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/location/v1"
	"testing"
)

var (
	locationClient v1.LocationClient
)

func init() {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		panic(err)
	}
	locationClient = v1.NewLocationClient(conn)
}

func TestLocationService_CreateLocation(t *testing.T) {
	reply, err := locationClient.CreateLocation(ctx, &v1.CreateLocationRequest{
		Type:    0,
		ReferID: 1463584877990383616,
		NodeID:  11111,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)
}

func TestLocationService_DeleteLocation(t *testing.T) {
	reply, err := locationClient.DeleteLocation(ctx, &v1.DeleteLocationRequest{
		Type:    0,
		ReferID: 1463584877990383616,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)
}

func TestLocationService_GetLocations(t *testing.T) {
	reply, err := locationClient.ListLocation(ctx, &v1.ListLocationRequest{
		Type:    0,
		ReferID: 1463584877990383616,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(reply)
}
