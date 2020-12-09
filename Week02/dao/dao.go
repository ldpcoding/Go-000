package dao

import (
	"github.com/pkg/errors"
	"week02/sql"
)

// 模拟 DAO 层获取数据
func FindData() (data interface{}, err error) {

	data, err = sql.Find()

	if err != nil {
		return nil, errors.Wrap(err, "dao: 获取数据出错！")
	}

	return data, err
}
