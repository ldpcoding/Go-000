package service

import "week02/biz"

// 模拟业务层获取数据
func GetData() (data interface{}, err error) {
	data, err = biz.GetData()
	if err != nil {
		return data, err
	}

	return data, err
}