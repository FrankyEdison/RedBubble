package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

/**
初始化雪花算法
param:开始时间，分布式机器id
*/
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime) //将字符串转换成time格式
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

// 获取一个雪花id
func GenerateID() int64 {
	return node.Generate().Int64()
}
