package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv31/cache"
	"github.com/liuhongdi/digv31/pkg/jwt"
	"github.com/liuhongdi/digv31/pkg/result"
	"github.com/liuhongdi/digv31/pkg/validCheck"
	"github.com/liuhongdi/digv31/request"
	"github.com/liuhongdi/digv31/service"
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
			//得到用户信息
			//用username生成origintoken
			originToken := jwt.GetOriginToken(param.UserName);
            //保存到cache
            cache.SetOneUserCache(originToken,userOne)
			//返回
			tokenString, _ := jwt.GenToken(originToken)

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
	originToken := c.MustGet("usertoken").(string)
	fmt.Println("user login begin")
	result := result.NewResult(c)
	userOne,err := cache.GetOneUserCache(originToken)
    if (err != nil) {
		result.Error(1,"需要登录")
	}else {
		result.Success(userOne)
	}
	return
}

//注销
func (u *UserController) Logout(c *gin.Context) {
	originToken := c.MustGet("usertoken").(string)
	fmt.Println("user login begin")
	result := result.NewResult(c)
	err := cache.DelOneUserCache(originToken)
	if (err != nil) {
		result.Error(1,err.Error())
	}else {
		result.Success("成功退出")
	}
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
	m := map[string]string {
		"origpassword":origpassword,
		"password":string(hashPwd),
	}
	result.Success(m)
	return
}
