// 模拟数据库驱动
package sql

import "errors"

var ErrNoRows = errors.New("sql: no rows")

// 模拟查询数据并抛出 ErrNoRows
func Find() (data interface{}, err error) {
	return nil, ErrNoRows
}
