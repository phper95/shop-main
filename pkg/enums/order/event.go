package order

const (
	Topic           = "shop-order"
	OperationCreate = "create"
	OperationUpdate = "update"
	OperationDelete = "delete"

	//-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：已完成；-1：已退款
)
