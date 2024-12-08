package models

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func InitSnowflakeID(machineID int64) {
	snowflakeNode, _ = snowflake.NewNode(machineID)
}

func NextSID() SID {
	return SID(snowflakeNode.Generate())
}

type SID snowflake.ID

func (i SID) String() string {
	return fmt.Sprintf("%d", i)
}

func ParseSID(in string) (SID, error) {
	id, err := snowflake.ParseString(in)
	return SID(id), err
}

func ParseSIDf(in string) SID {
	id, _ := snowflake.ParseString(in)
	return SID(id)
}
