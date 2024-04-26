package server

import (
	v12 "master/api/bucket/v1"
	v1 "master/api/user/v1"
	"master/internal/conf"
	"master/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	bucketService *service.BucketService,
	nodeService *service.NodeService,
	objectService *service.ObjectService,
	blockService *service.BlockService,
	aclService *service.AclService,
	userService *service.UserService,
	locationService *service.LocationService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, userService)
	v12.RegisterBucketHTTPServer(srv, bucketService)
	return srv
}
