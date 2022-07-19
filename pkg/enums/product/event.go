package product

const (
	//手动创建topic
	// bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 -- topic shop-product
	Topic           = "shop-product"
	OperationCreate = "create"
	OperationUpdate = "update"
	OperationDelete = "delete"
	OperationOnSale = "on_sale"
	OperationUnSale = "un_sale"
)
