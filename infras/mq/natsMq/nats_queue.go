package natsMq

import "github.com/nats-io/nats.go"

/*
基于队列组的主题订阅：
具有相同队列名称的所有订阅都将形成一个队列组。使用队列语义，每个消息将仅传递给每个队列组的一个订阅服务器。
您可以拥有任意数量的队列组。普通订阅服务器将继续按预期工作。
*/
func QueueSubscribe(subject,queue string, handler nats.Handler) error {
	conn, err := NatsMqConnPool.Get()
	if err != nil {
		return err
	}
	defer NatsMqConnPool.Put(conn)

	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	_, err = encodedConn.QueueSubscribe(subject, queue, handler)
	return err
}
