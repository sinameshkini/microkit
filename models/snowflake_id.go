package models

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func InitSnowflakeID(machineID int64) {
	snowflakeNode, _ = snowflake.NewNode(machineID)
}

func Next() SID {
	return SID(snowflakeNode.Generate())
}

type SID snowflake.ID

func (i SID) String() string {
	return fmt.Sprintf("%d", i)
}

func ParseID(in string) (SID, error) {
	id, err := snowflake.ParseString(in)
	return SID(id), err
}

func ParseIDf(in string) SID {
	id, _ := snowflake.ParseString(in)
	return SID(id)
}
