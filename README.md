# 海量数据高并发场景，构建Go+ES8企业级搜索微服务课程商城系统

## 项目简介
海量数据高并发场景，构建Go+ES8企业级搜索微服务课程商城系统基于当前流行技术组合的前后端商城管理系统：

[课程地址 **点此 打开**](https://coding.imooc.com/class/579.html?mc_marking=bb86c9071ed9b7cf12612a2a85203372)

项目基于Gin+Gorm+Casbin+Jwt+Redis+Mysql5.7+Vue 的前后端分离电商管理系统，权限控制采用RBAC，
支持商城商品加入购物车、下单、评价、支付（微信支付与余额支付）、搜索、地址管理、快递鸟查询、多级分类，商品管理、商品sku、图片素材、数据字典与数据权限管理，支持动态路由等

## pc商城功能：

- 登录注册
- 首页轮播图
- 首页商品展示
- 商品详情及其sku
- 商品加入购物车
- 商品下单
- 商品支付（微信与余额支付）
- 商品个人订单管理
- 商品的收藏
- 商品的地址管理
- 商品的评价管理
- 快递鸟快递查询
- 商品分类等搜索
- 个人中心图像上传等

##  商城后台系统功能

- 用户管理：提供用户的相关配置 
- 角色管理：对权限与菜单进行分配，可根据部门设置角色的数据权限 
- 菜单管理：已实现菜单动态路由，后端可配置化 
- 部门管理：可配置系统组织架构，树形表格展示 
- 岗位管理：配置各个部门的职位 
- 字典管理：可维护常用一些固定的数据，如：状态，性别等 
- 日志管理：用户操日志记录 
- 素材管理：图片素材库 <br>
- 分类管理：商品多级分类 <br>
- sku管理：商品sku规则管理 <br>
- 商品管理：可以添加单规格或者多规格商品含有百度编辑器 <br>
- 微信公众号：可微信图文、微信菜单等 <br>
- 订单管理：对订单发货查看详情等操作
- 物流快递：实现了快递鸟基本查询功能

## 项目结构

```
- internal 应用代码
    - controllers 控制器模块
      - admin 后端控制器
      - front 前端控制器
    - listen redis监听器
    - models 模型模块
    - service 服务模块
      - product_serive 商品服务
      - wechat_menu_serive 微信公众号菜单服务
      ......
- conf 公共配置
  -config.yml yml配置文件
  -config.go 配置解析，转化成对应的结构体
  
- middleware 中间件
    - AuthCheck.go  jwt接口权限校验
	- cors.go 跨域处理
	......
- pkg 程序应用包
  - app
  - base
  - casbin
  - jwt
  - qrcode
  - wechat
  .....
- routere 路由
- logs 日志存放
- runtime 资源目录
```

## 环境要求
- go >= 1.15
- MySQL >= 5.7
- redis >=4.0.0

## 后端技术
gin、gorm、jwt、redis、Mysql、copier、ksuid、 Redis、zap、viper、wechat

##　前端技术
npm、ES6、vue-cli、vue-router、vuex、element-ui

## 后端配置部署和启动

```
1、安装go>=1.15,这个可以https://studygolang.com/dl下载

2、开启mod： go env -w GO111MODULE=on

3、配置代理：go env -w GOPROXY=https://goproxy.cn,direct 这个让下载依赖速度更快

5、配置私有仓库：go env -w  GOPRIVATE=*gitee.com

6、下载项目：git clone https://gitee.com/phper95/shop-main.git

7、go mod tidy 安装所需依赖

8、导入sql/shop.sql,修改cconfig,yml 里数据库与redis配置

9、本地运行go run main.go

```

##  线上部署：

## 权限检验说明

```
//权限校验中间件路径./middleware/authcheck.go 里面 
//注意下面注释的代码块，此处用于项目演示，请注意删除

func Jwt() gin.HandlerFunc {
return func(c *gin.Context) {
var data interface{}
var appG = app.Gin{C: c}

      url := c.Request.URL.Path

      method := strings.ToLower(c.Request.Method)
      //部署线上开启
      //prohibit := "post,put,delete"
      //if url != "/admin/auth/logout" && strings.Contains(prohibit,method) {
      // ctx.Output.JSON(controllers.ErrMsg("演示环境禁止操作",40006),
      //    true,true)
      // return
      //}

      mytoken := c.Request.Header.Get("Authorization")
      if len(mytoken) < bearerLength {
         appG.Response(http.StatusUnauthorized,constant.ERROR_AUTH,data)
         c.Abort()
         return
      }
```



## 后端技术选型
* gin
* jwt
* redis
* Mysql8
* Gorm
* copier
* ksuid
* Redis
*  Casbin
*  viper
*  zap
*  wecchat
*  gopay
## 前端技术选型
* npm
* ES6
* vue-cli
* vue-router
* vuex
* element-ui 

## 账号密码
前台账号密码:
后台账号密码: admin/123456


## 免责声明
商城代码仅用于学习演示，如果需要用于线上环境请严格测试，本人不承担代码问题带来的任何损失
