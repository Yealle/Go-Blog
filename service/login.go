package service

import (
	"Blog/dao"
	"Blog/utils"
	"errors"
	"models"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	// 加密
	passwd = utils.Md5Crypt(passwd, "216")
	// fmt.Println(passwd)

	user := dao.GetUser(userName, passwd)

	if user == nil {
		return nil, errors.New("账号密码不正确")
	}

	uid := user.Uid

	// 生成token jwt技术生成令牌
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}

	var userInfo models.UserInfo

	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var loginRes = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}

	return loginRes, nil
}
