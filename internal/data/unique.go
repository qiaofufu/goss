package data

import (
	"master/internal/conf"
	"master/third_party/snowflake"
)

type UniqueIDGenerator interface {
	GenerateID() int64
}

func NewUniqueIDGenerator(node *conf.Node) UniqueIDGenerator {
	return snowflake.NewSnowflake(node.Id, node.StartTime)
}
