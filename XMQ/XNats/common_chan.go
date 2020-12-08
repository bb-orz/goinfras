package XNats

import (
	"github.com/nats-io/nats.go"
	"goinfras"
)

type commonNatsChan struct {
	pool *NatsPool
}

// 发送消息到一个主题，绑定管道
func (c *commonNatsChan) BindSendChan(subject string, sendCh chan interface{}) error {
	conn, err := c.pool.Get()
	if err != nil {
		return err
	}
	defer c.pool.Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	err = encodedConn.BindSendChan(subject, sendCh)
	return err
}

// 接收主题消息，绑定管道
func (c *commonNatsChan) BindRecvChan(subject string, recvCh chan interface{}) error {
	conn, err := c.pool.Get()
	if err != nil {
		return err
	}
	defer c.pool.Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	_, err = encodedConn.BindRecvChan(subject, recvCh)
	return err
}
