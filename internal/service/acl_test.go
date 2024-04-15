package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/acl/v1"
	"testing"
)

var (
	aclClient v1.AclClient
)

func init() {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		panic(err)
	}
	aclClient = v1.NewAclClient(conn)
}

func TestAclService_Create(t *testing.T) {
	acl, err := aclClient.CreateAcl(ctx, &v1.CreateAclRequest{
		ResourceId:   3431232960148279296,
		ResourceType: 0,
		TargetId:     1001,
		TargetType:   0,
		Permission:   1,
		OperatorId:   1001,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(acl)
}

func TestAclService_Update(t *testing.T) {
	acl, err := aclClient.UpdateAcl(ctx, &v1.UpdateAclRequest{
		Id:         5010713583265255424,
		Permission: 2,
		OperatorId: 1001,
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(acl)
}

func TestAclService_Delete(t *testing.T) {
	acl, err := aclClient.DeleteAcl(ctx, &v1.DeleteAclRequest{
		Id: 5010713583265255424,
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(acl)
}

func TestAclService_List(t *testing.T) {
	acl, err := aclClient.ListAcl(ctx, &v1.ListAclRequest{
		ResourceId: 3431232960148279296,
		OperatorId: 1001,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(acl)
}
