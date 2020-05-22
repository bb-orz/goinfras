package natsMq

import "github.com/nats-io/nats.go"

// 发送消息到一个主题，绑定管道
func BindSendChan(subject string,sendCh chan interface{}) error {
	conn, err := NatsMqConnPool.Get()
	if err != nil {
		return err
	}
	defer NatsMqConnPool.Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	err = encodedConn.BindSendChan(subject, sendCh)
	return err
}

// 接收主题消息，绑定管道
func BindRecvChan(subject string,recvCh chan interface{}) error {
	conn, err := NatsMqConnPool.Get()
	if err != nil {
		return err
	}
	defer NatsMqConnPool.Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	_, err = encodedConn.BindRecvChan(subject, recvCh)
	return err
}

// 基于队列的接收操作，绑定通道。
func BindRecvQueueChan(subject,queue string,recvCh chan interface{}) error {
	conn, err := NatsMqConnPool.Get()
	if err != nil {
		return err
	}
	defer NatsMqConnPool.Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	_, err = encodedConn.BindRecvQueueChan(subject,queue, recvCh)
	return err
}