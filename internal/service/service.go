package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBucketService, NewNodeService, NewObjectService, NewBlockService, NewAclService, NewUserService, NewLocationService)
