package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"time"
)

// 存放业务逻辑的代码

//用户注册
func SignUp(p *models.ParamSignUp) (err error) {

	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//生成idTODO,雪花算法生成id
	var userID int64
	userID = 1
	// 构造一个User实例
	user := &models.User{
		UserID:     userID,
		Username:   p.Username,
		Password:   p.Password,
		CreateTime: time.Now().Unix(),
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}
