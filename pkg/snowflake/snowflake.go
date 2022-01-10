package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2022-01-07", 1); err != nil {
//		fmt.Printf("init failed, err:%v\n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}
