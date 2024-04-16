package server

import (
	aclAPI "master/api/acl/v1"
	blockAPI "master/api/block/v1"
	bucketAPI "master/api/bucket/v1"
	nodeAPI "master/api/node/v1"
	objectAPI "master/api/object/v1"
	userAPI "master/api/user/v1"
	"master/internal/conf"
	"master/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	bucket *service.BucketService,
	node *service.NodeService,
	object *service.ObjectService,
	block *service.BlockService,
	aclService *service.AclService,
	userService *service.UserService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	bucketAPI.RegisterBucketServer(srv, bucket)
	nodeAPI.RegisterNodeServer(srv, node)
	objectAPI.RegisterObjectServer(srv, object)
	blockAPI.RegisterBlockServer(srv, block)
	aclAPI.RegisterAclServer(srv, aclService)
	userAPI.RegisterUserServer(srv, userService)
	return srv
}
