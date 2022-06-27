package front

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"shop/internal/service/wechat_user_service"
	"shop/pkg/global"
	"shop/pkg/util"
)

// 公众号服务api
type WechatController struct {
}

// @Title 公众号服务
// @Description 公众号服务
// @Success 200 {object} app.Response
// @router / [any]
func (e *WechatController) GetAll(c *gin.Context) {
	official := global.WechatOfficial
	server := official.GetServer(c.Request, c.Writer)

	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		if msg.MsgType == message.MsgTypeEvent {
			global.LOG.Info(msg.Event)
			if msg.Event == message.EventSubscribe {
				//存储用户
				user := official.GetUser()
				userInfo, e := user.GetUserInfo(msg.CommonToken.GetOpenID())
				if e != nil {
					global.LOG.Error(e)
				}
				ip := util.GetClientIP(c)
				userSerive := wechat_user_service.User{UserInfo: userInfo, Ip: ip}
				userSerive.Insert()
			}
		}
		global.LOG.Info(msg.MsgType)
		text := message.NewText(msg.Content)

		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		global.LOG.Error(err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		global.LOG.Error(err)
		return
	}

}
