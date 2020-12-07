package XNats

import "github.com/nats-io/nats.go"

type commonNatsQueue struct {
	pool *NatsPool
}

/*
基于队列组(Queue)的主题订阅：
具有相同队列名称的所有订阅都将形成一个队列组。使用队列语义，每个消息将仅传递给每个队列组的一个订阅服务器。
您可以拥有任意数量的队列组。普通订阅服务器将继续按预期工作。

说明：Queue=工作组，工作组中有N个worker，发布消息后，同一个工作组中，仅有一个worker会收到消息。


// Handlers 接收以下四种签名之一，接收定制的消息处理.
//
//	type person struct {
//		Name string `json:"name,omitempty"`
//		Age  uint   `json:"age,omitempty"`
//	}
//
//	handler := func(m *Msg)
//	handler := func(p *person)
//	handler := func(subject string, o *obj)
//	handler := func(subject, reply string, o *obj)
*/
func (c *commonNatsQueue) QueueSubscribe(subject, queue string, handler nats.Handler) error {
	conn, err := c.pool.Get()
	if err != nil {
		return err
	}
	defer c.pool.Put(conn)

	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	_, err = encodedConn.QueueSubscribe(subject, queue, handler)
	return err
}

// 基于队列组的接收操作，绑定通道。
func (c *commonNatsQueue) BindRecvQueueChan(subject, queue string, recvCh chan interface{}) error {
	conn, err := c.pool.Get()
	if err != nil {
		return err
	}
	defer c.pool.Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	_, err = encodedConn.BindRecvQueueChan(subject, queue, recvCh)
	return err
}
