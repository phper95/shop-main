package pay_service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/wechat"
	"github.com/unknwon/com"
	"gorm.io/gorm"
	"shop/internal/models"
	"shop/internal/service/order_service"
	orderDto "shop/internal/service/order_service/dto"
	"shop/internal/service/wechat_user_service"
	orderEnum "shop/pkg/enums/order"
	"shop/pkg/global"
	"time"
)

type Pay struct{}

//开始支付
func GoPay(returnMap map[string]interface{}, orderId, payType, from string,
	uid int64, dto *orderDto.OrderExtend) (map[string]interface{}, *models.StoreOrder, error) {
	var order *models.StoreOrder
	dto = &orderDto.OrderExtend{}

	switch payType {
	case "weixin":
		if from == "pc" {
			client := wechat.NewClient("", "", "", true)

			//设置国家
			client.SetCountry(wechat.China)
			client.DebugSwitch = gopay.DebugOn
			orderService := order_service.Order{
				Uid:     uid,
				OrderId: orderId,
			}
			orderInfo, order, _ := orderService.GetOrderInfo()
			//expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
			bm := make(gopay.BodyMap)
			bm.Set("nonce_str", util.RandomString(32)).
				Set("body", "shop-go pc支付").
				Set("out_trade_no", orderId).
				Set("total_fee", orderInfo.PayPrice*100).
				Set("spbill_create_ip", "127.0.0.1").
				Set("notify_url", global.CONFIG.App.Domain+"/order/notify").
				Set("trade_type", wechat.TradeType_Native).
				Set("device_info", "WEB").
				Set("sign_type", wechat.SignType_MD5)
			var ctx = context.Background()
			wxRsp, err := client.UnifiedOrder(ctx, bm)
			if err != nil {
				global.LOG.Error(err)
				return nil, order, err
			}
			global.LOG.Info(wxRsp)

			jsConfig := gin.H{"code_url": wxRsp.CodeUrl}
			dto.JsConfig = jsConfig
			returnMap["payMsg"] = "pc支付成功"
			returnMap["result"] = dto
			orderService.M = order
		}
	case "yue":
		order, err := YuePay(orderId, uid)
		if err != nil {
			global.LOG.Error(err)
			return nil, order, err
		}
		returnMap["payMsg"] = "余额支付成功"

	}

	return returnMap, order, nil
}

//余额支付
func YuePay(orderId string, uid int64) (*models.StoreOrder, error) {
	var err error
	orderService := order_service.Order{
		Uid:     uid,
		OrderId: orderId,
	}
	orderInfo, order, _ := orderService.GetOrderInfo()
	if orderInfo.Paid == 1 {
		return order, errors.New("订单已经支付")
	}

	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userService := wechat_user_service.User{
		Id: uid,
	}
	userInfo := userService.GetUserInfo()
	global.LOG.Info(userInfo.NowMoney, orderInfo.PayPrice)
	if userInfo.NowMoney < orderInfo.PayPrice {
		return order, errors.New("余额不足")
	}
	err = tx.Exec("update user set now_money=now_money - ?"+
		" where id = ?", orderInfo.PayPrice, uid).Error
	if err != nil {
		global.LOG.Error(err)
		return order, errors.New("余额支付失败-0001")
	}
	order, err = PaySuccess(tx, orderInfo.OrderId, "yue")
	if err != nil {
		global.LOG.Error(err)
		return order, errors.New("余额支付失败-0002")
	}
	return order, nil
}

//支付成功处理
func PaySuccess(tx *gorm.DB, orderId, payType string) (*models.StoreOrder, error) {
	var err error

	orderService := order_service.Order{
		OrderId: orderId,
	}
	orderInfo, order, _ := orderService.GetOrderInfo()

	//修改订单状态
	updateOrder := &models.StoreOrder{
		Paid:    orderEnum.PAY_STATUS_1,
		PayType: payType,
		PayTime: time.Now(),
	}
	err = tx.Model(&models.StoreOrder{}).Where("order_id = ?", orderId).Updates(updateOrder).Error

	if err != nil {
		global.LOG.Error(err)
		return order, err
	}
	order.Paid = updateOrder.Paid
	order.PayType = updateOrder.PayType
	order.PayTime = updateOrder.PayTime

	//增加用户购买次数
	err = tx.Exec("update user set pay_count = pay_count + 1"+
		" where id = ?", orderInfo.Uid).Error
	if err != nil {
		global.LOG.Error(err)
		return order, err
	}
	//增加状态
	err = models.AddStoreOrderStatus(tx, orderInfo.Id, "pay_success", "用户付款成功")
	if err != nil {
		global.LOG.Error(err)
		return order, err
	}

	userServie := wechat_user_service.User{Id: orderInfo.Uid}
	userInfo := userServie.GetUserInfo()
	payTypeMsg := "微信支付"
	if payType == "yue" {
		payTypeMsg = "余额支付"
	}

	mark := payTypeMsg + com.ToStr(orderInfo.PayPrice) + "元购买商品"
	err = models.Expend(tx, orderInfo.Uid, "购买商品", "now_money", "pay_product",
		mark, com.ToStr(orderInfo.Id), orderInfo.PayPrice, userInfo.NowMoney)
	if err != nil {
		global.LOG.Error(err)
		return order, err
	}

	return order, nil
	//todo 消息通知
}
