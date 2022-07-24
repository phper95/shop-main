package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/phper95/pkg/cache"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"shop/internal/models"
	"shop/internal/models/vo"
	"shop/pkg/constant"
	"shop/pkg/global"
	"shop/pkg/logging"
	"strconv"
	"strings"
	"time"
)

var jwtSecret []byte

const bearerLength = len("Bearer ")

var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

type userStdClaims struct {
	vo.JwtUser
	//*models.User
	jwt.StandardClaims
}

func Init() {
	jwtSecret = []byte(global.CONFIG.App.JwtSecret)
}

func GenerateAppToken(m *models.ShopUser, d time.Duration) (string, error) {
	m.Password = ""
	//m.Permissions = []string{}
	//expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(d).Unix(),
		Id:        strconv.FormatInt(m.Id, 10),
		Issuer:    "shopAppGo",
	}

	var jwtUser = vo.JwtUser{
		Id:       m.Id,
		Avatar:   m.Avatar,
		Username: m.Username,
		Phone:    m.Phone,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		JwtUser:        jwtUser,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logging.Error(err)
	}
	//set redis
	var key = constant.AppRedisPrefixAuth + tokenString
	json, _ := json.Marshal(m)
	err = cache.GetRedisClient(cache.DefaultRedisClient).Set(key, json, d)
	if err != nil {
		global.LOG.Error("GenerateAppToken cache set error", err, "key", key)
	}

	return tokenString, err
}

//返回id
func GetAppUserId(c *gin.Context) (int64, error) {
	u, exist := c.Get(constant.AppAuthUser)
	if !exist {
		shopUser, err := GetAppDetailUser(c)
		if err != nil {
			return 0, err
		}
		return shopUser.Id, nil
	}
	user, ok := u.(*vo.JwtUser)

	if ok {
		return user.Id, nil
	}
	return 0, errors.New("can't convert to user struct")
}

//返回user
func GetAppUser(c *gin.Context) (*vo.JwtUser, error) {
	u, exist := c.Get(constant.AppAuthUser)
	if !exist {
		shopUser, err := GetAppDetailUser(c)
		if err != nil {
			return nil, err
		}
		jwtUser := vo.JwtUser{
			Id:       shopUser.Id,
			Avatar:   shopUser.Avatar,
			Username: shopUser.Username,
			Phone:    shopUser.Phone,
			NickName: shopUser.Nickname,
		}
		return &jwtUser, nil
	}
	user, ok := u.(*vo.JwtUser)
	if ok {
		return user, nil
	}
	return nil, errors.New("can't convert to user struct")
}

//返回 detail user
func GetAppDetailUser(c *gin.Context) (*models.ShopUser, error) {
	mytoken := c.Request.Header.Get("Authorization")
	if mytoken == "" {
		return nil, errors.New("user not login")
	}
	token := strings.TrimSpace(mytoken[bearerLength:])
	var key = constant.AppRedisPrefixAuth + token
	val, err := cache.GetRedisClient(cache.DefaultRedisClient).GetStr(key)
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]string)
	userMap[key] = val
	jsonStr := userMap[key]
	user := &models.ShopUser{}
	err = json.Unmarshal([]byte(jsonStr), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func RemoveAppUser(c *gin.Context) error {
	mytoken := c.Request.Header.Get("Authorization")
	token := strings.TrimSpace(mytoken[bearerLength:])
	var key = constant.AppRedisPrefixAuth + token
	return cache.GetRedisClient(cache.DefaultRedisClient).Delete(key)
}

func GenerateToken(m *models.SysUser, d time.Duration) (string, error) {
	m.Password = ""
	stdClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(d).Unix(),
		Id:        strconv.FormatInt(m.Id, 10),
		Issuer:    "shopGo",
	}

	var (
		roleNames []string
	)
	for _, role := range m.Roles {
		roleNames = append(roleNames, role.Permission)
	}

	var jwtUser = vo.JwtUser{
		Id:       m.Id,
		Avatar:   m.Avatar,
		Email:    m.Email,
		Username: m.Username,
		Phone:    m.Phone,
		NickName: m.NickName,
		Sex:      m.Sex,
		Dept:     m.Depts.Name,
		Job:      m.Jobs.Name,
		Roles:    roleNames,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		JwtUser:        jwtUser,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logging.Error(err)
	}
	//set redis
	var key = constant.RedisPrefixAuth + tokenString
	json, _ := json.Marshal(m)
	cache.GetRedisClient(cache.DefaultRedisClient).Set(key, string(json), d)

	return tokenString, err
}

func ValidateToken(tokenString string) (*vo.JwtUser, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	return &claims.JwtUser, err

}

//返回id
func GetAdminUserId(c *gin.Context) (int64, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("can't get user id")
	}
	user, ok := u.(*vo.JwtUser)

	if ok {
		return user.Id, nil
	}
	return 0, errors.New("can't convert to user struct")
}

//返回user
func GetAdminUser(c *gin.Context) (*vo.JwtUser, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get user id")
	}
	user, ok := u.(*vo.JwtUser)
	if ok {
		return user, nil
	}
	return nil, errors.New("can't convert to user struct")
}

//返回 detail user
func GetAdminDetailUser(c *gin.Context) *models.SysUser {
	mytoken := c.Request.Header.Get("Authorization")
	token := strings.TrimSpace(mytoken[bearerLength:])
	var key = constant.RedisPrefixAuth + token
	val, err := cache.GetRedisClient(cache.DefaultRedisClient).GetStr(key)
	if err != nil {
		global.LOG.Error("redis error ", err, "key", key, "cmd : Get", "client", cache.DefaultRedisClient)
		return nil
	}
	userMap := make(map[string]string)
	userMap[key] = val
	jsonStr := userMap[key]
	user := &models.SysUser{}
	json.Unmarshal([]byte(jsonStr), user)
	return user
}

func RemoveUser(c *gin.Context) error {
	mytoken := c.Request.Header.Get("Authorization")
	token := strings.TrimSpace(mytoken[bearerLength:])
	var key = constant.RedisPrefixAuth + token
	return cache.GetRedisClient(cache.DefaultRedisClient).Delete(key)
}
