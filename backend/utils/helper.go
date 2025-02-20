package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// TODO: need to set at env for NewNode value
func GenerateSnowflakeID() uint64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	id := node.Generate().Int64()
	return uint64(id)
}