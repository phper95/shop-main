package wechat_user_service

import (
	"encoding/json"
	"errors"
	"gitee.com/phper95/pkg/cache"
	"github.com/jinzhu/copier"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"shop/internal/models"
	"shop/internal/models/vo"
	"shop/internal/params"
	userDto "shop/internal/service/user_service/dto"
	wechatUserDto "shop/internal/service/wechat_user_service/dto"
	wechatUserVo "shop/internal/service/wechat_user_service/vo"
	"shop/pkg/constant"
	userEnum "shop/pkg/enums/user"
	"shop/pkg/global"
	"shop/pkg/util"
	"time"
)

type User struct {
	Id       int64
	Username string

	Value    string
	MyType   string
	UserType string

	PageNum  int
	PageSize int

	UserInfo *user.Info

	Ip string
	//M *models.ShopUser
	Dto *wechatUserDto.ShopUser

	Money *userDto.UserMoney

	Ids []int64

	HLoginParam *params.HLoginParam
	RegParam    *params.RegParam
	VerityParam *params.VerityParam

	User *models.ShopUser
}

func (u *User) GetUserInfo() *wechatUserVo.User {
	var (
		userVO wechatUserVo.User
		user   models.ShopUser
	)
	global.Db.Model(&models.ShopUser{}).Where("id = ?", u.Id).First(&user)
	copier.Copy(&userVO, user)

	return &userVO
}

func (u *User) GetUserDetail() *wechatUserVo.User {
	var user wechatUserVo.User

	e := copier.Copy(&user, u.User)
	if e != nil {
		global.LOG.Error(e)
	}

	return &user
}

func (u *User) Reg() error {
	var (
		user models.ShopUser
		err  error
	)
	err = global.Db.
		Model(&models.ShopUser{}).
		Where("username = ?", u.RegParam.Account).First(&user).Error
	if err == nil {
		return errors.New("用户已经存在")
	}
	codeKey := constant.SmsCode + u.RegParam.Account
	code, err := cache.GetRedisClient(cache.DefaultRedisClient).GetStr(codeKey)
	if err != nil {
		global.LOG.Error("redis error", "key : ", codeKey, "cmd : get")
	}
	if code != u.RegParam.Captcha {
		return errors.New("验证码不对")
	}

	uu := models.ShopUser{
		Username: u.RegParam.Account,
		Nickname: u.RegParam.Account,
		Password: util.HashAndSalt([]byte(u.RegParam.Password)),
		RealName: u.RegParam.Account,
		Avatar:   "",
		AddIp:    u.Ip,
		LastIp:   u.Ip,
		UserType: userEnum.PC,
		Phone:    u.RegParam.Account,
	}
	err = models.AddWechatUser(&uu)
	//注册成功删除验证码缓存
	cache.GetRedisClient(cache.DefaultRedisClient).Delete(code)
	return err

}

func (u *User) Verify() (string, error) {
	var (
		user models.ShopUser
		err  error
	)
	err = global.Db.
		Model(&models.ShopUser{}).
		Where("username = ?", u.VerityParam.Phone).First(&user).Error
	if err == nil {
		return "", errors.New("手机已经注册过")
	}

	codeKey := constant.SmsCode + u.VerityParam.Phone
	exists, err := cache.GetRedisClient(cache.DefaultRedisClient).Exists(codeKey)
	if err != nil {
		global.LOG.Error("redis error ", err, "key", codeKey, "cmd : Exists", "client", cache.DefaultRedisClient)
	}
	codeVal, err := cache.GetRedisClient(cache.DefaultRedisClient).GetStr(codeKey)
	if err != nil {
		global.LOG.Error("redis error ", err, "key", codeKey, "cmd : Get", "client", cache.DefaultRedisClient)
	}
	if exists {
		return "", errors.New("10分钟有效:" + codeVal)
	}

	code := util.RandomNumber(constant.SmsLength)
	expireTime := time.Minute * 10
	err = cache.GetRedisClient(cache.DefaultRedisClient).Set(codeKey, code, expireTime)
	if err != nil {
		global.LOG.Error("redis error ", err, "key", codeKey, "cmd : Set", "client", cache.DefaultRedisClient)
	}
	//此处发送阿里云短信
	//测试阶段直接把验证码返回
	return "测试阶段验证码为：" + code, nil

}

func (u *User) GetUserAll() vo.ResultList {
	maps := make(map[string]interface{})

	if u.Value != "" {
		if u.MyType == "phone" {
			maps["phone"] = u.Value
		} else {
			maps["nickname"] = u.Value
		}
	}

	if u.UserType != "" {
		maps["user_type"] = u.UserType
	}

	total, list := models.GetAllWechatUser(u.PageNum, u.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (u *User) Insert() error {
	result, _ := json.Marshal(u.UserInfo)
	user := models.ShopUser{
		Username:  u.UserInfo.OpenID,
		Nickname:  u.UserInfo.Nickname,
		Password:  util.HashAndSalt([]byte("123456")),
		RealName:  u.UserInfo.Nickname,
		Avatar:    u.UserInfo.Headimgurl,
		AddIp:     u.Ip,
		LastIp:    u.Ip,
		UserType:  userEnum.WECHAT,
		WxProfile: datatypes.JSON(result),
	}
	return models.AddWechatUser(&user)
}

func (u *User) Save() error {
	user := models.ShopUser{
		RealName: u.Dto.RealName,
		Mark:     u.Dto.Mark,
		Phone:    u.Dto.Phone,
		Integral: u.Dto.Integral,
	}
	return models.UpdateByWechatUsere(u.Dto.Id, &user)
}

func (u *User) SaveMony() error {
	var err error
	if u.Money.Ptype == 1 {
		err = global.Db.
			Model(&models.ShopUser{}).
			Where("id = ?", u.Money.Id).
			Update("now_money", gorm.Expr("now_money + ?", u.Money.Money)).Error
	} else {
		err = global.Db.
			Model(&models.ShopUser{}).
			Where("id = ? and now_money >= ?", u.Money.Id, u.Money.Money).
			Update("now_money", gorm.Expr("now_money - ?", u.Money.Money)).Error
	}
	return err
}

func (u *User) HLogin() (*models.ShopUser, error) {
	var (
		user models.ShopUser
		err  error
	)
	err = global.Db.
		Model(&models.ShopUser{}).
		Where("username = ?", u.HLoginParam.Username).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if !util.ComparePwd(user.Password, []byte(u.HLoginParam.Password)) {
		return nil, errors.New("密码不对")
	}

	return &user, err

}
