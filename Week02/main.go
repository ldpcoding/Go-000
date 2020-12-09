package main

import (
	"fmt"
	"week02/service"
)

func main() {
	// 模拟获取数据
	data, err := service.GetData()
	if err != nil {
		fmt.Println(err) // 模拟打印日志

		return
	}

	fmt.Println(data) // 模拟响应数据
}
