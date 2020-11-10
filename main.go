package main

import (
	"bluebell/settings"
	"fmt"
)

func main() {
	//1、加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
}
