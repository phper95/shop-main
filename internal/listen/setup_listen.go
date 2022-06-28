package listen

import (
	"fmt"
	"shop/pkg/global"
)

func Init() {
	var sub PSubscriber
	fmt.Printf(global.CONFIG.Redis.Host)
	conn := PConnect(global.CONFIG.Redis.Host, global.CONFIG.Redis.Password)
	sub.ReceiveKeySpace(conn)
	sub.Psubscribe()
}
