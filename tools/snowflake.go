// @Title tools
// @Description // TODO
// @Author chenhaoran
// @Datetime  2021/2/24 11:21 上午
package tools

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitSnowflake(int int64) error {

	// Create snowflake node
	n, err := snowflake.NewNode(int)
	if err != nil {
		return err
	}
	// Set node
	node = n
	return nil
}

// GenerateSnowflake generates Twitter Snowflake ID
func GenerateSnowflake() int64 {
	return int64(node.Generate())
}
