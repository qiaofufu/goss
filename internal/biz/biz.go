package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBucketUsecase, NewConsistentHash, NewNodeUsecase, NewObjectUsecase, NewBlockUsecase, NewAclUsecase)
