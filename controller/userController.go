package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv07/pkg/jwt"
	"github.com/liuhongdi/digv07/pkg/result"
	"github.com/liuhongdi/digv07/pkg/validCheck"
	"github.com/liuhongdi/digv07/request"
	"github.com/liuhongdi/digv07/service"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}
//用户登录
func (u *UserController) Login(c *gin.Context) {
	fmt.Println("user login begin")
	result := result.NewResult(c)
	param := request.LoginRequest{
		UserName: c.Param("username"),
		PassWord: c.Param("password"),
		}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400,errs.Error())
		return
	}

    userOne,err := service.GetOneUser(param.UserName)
    if err!=nil {
		result.Error(404,"数据查询错误")
	} else {
        //password is right?
		err := bcrypt.CompareHashAndPassword([]byte(userOne.Password), []byte(param.PassWord))
		// 没有错误则密码匹配
		if err != nil {
			result.Error(1001,"账号信息错误")
			//return false
		}else {
			//生成token并返回
			fmt.Println("begin gentoken:")
			fmt.Println(param.UserName)

			tokenString, _ := jwt.GenToken(param.UserName)

			m := map[string]string {
				"tokenString":tokenString,
			}
            result.Success(m)
			//return
		}
	}
	return
}

//用户信息info
func (u *UserController) Info(c *gin.Context) {
	username := c.MustGet("username").(string)
	fmt.Println("user login begin")
	result := result.NewResult(c)

	m := map[string]string {
		"username":username,
	}
	result.Success(m)

	return
}

//生成用户pass
func (u *UserController) Pass(c *gin.Context) {
	//origpassword := c.MustGet("password").(string)
	origpassword := c.Query("password")
	fmt.Println("password:"+origpassword)
	fmt.Println("user pass begin")
	result := result.NewResult(c)

	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(origpassword), bcrypt.DefaultCost)

	/*
	//compare
	err := bcrypt.CompareHashAndPassword(hashPwd, []byte("124"))
	parse:=""
	if (err == nil) {
		parse="true"
	} else {
		parse = "false"
	}
	// 没有错误则密码匹配
	if err != nil {
		log.Println(err)
	}
    */
	m := map[string]string {
		"origpassword":origpassword,
		"password":string(hashPwd),
		//"parse":parse,
	}
	result.Success(m)

	return
}
