package dao

import (
"fmt"
"github.com/liuhongdi/digv31/global"
"github.com/liuhongdi/digv31/model"
)
//select一条记录
func SelectOneUser(userName string) (*model.User, error) {
	fields := []string{"userId", "username", "password","addTime"}
	userOne:=&model.User{}
	err := global.DBLink.Select(fields).Where("username=?",userName).First(&userOne).Error
	if (err != nil) {
		return nil,err
	} else {
		fmt.Println(userOne)
		return userOne,nil
	}
}
