package service

import (
"github.com/liuhongdi/digv31/dao"
"github.com/liuhongdi/digv31/model"
)
//得到一篇文章的详情
func GetOneUser(userName string) (*model.User, error) {
	return dao.SelectOneUser(userName)
}
