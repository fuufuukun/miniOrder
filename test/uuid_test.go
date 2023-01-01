package test

import (
	"fmt"
	"miniorder/util"
	"testing"
)

func TestUUIDGenerator(t *testing.T) {
	//新建UUIDgenerator
	UUIDFactory := util.NewUUIDGenerator("idtest")

	//获取UUID
	for i := 0; i < 50; i++ {
		fmt.Println(UUIDFactory.Get())
	}

	//获取uint32形式的uuid
	for i := 0; i < 50; i++ {
		fmt.Println(UUIDFactory.GetUint32())
	}
}
