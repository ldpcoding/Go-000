package biz

import (
	"week02/dao"
)

// 模拟业务层获取数据
func GetData() (data interface{}, err error) {
	data, err = dao.FindData()
	if err != nil {
		return data, err
	}

	return data, err
}
