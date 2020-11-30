package natsMq

import "github.com/nats-io/nats.go"

type CommonNatsChan struct{}

func NewCommonNatsChan() *CommonNatsChan {
	return new(CommonNatsChan)
}

// 发送消息到一个主题，绑定管道
func (*CommonNatsChan) BindSendChan(subject string, sendCh chan interface{}) error {
	conn, err := NatsMQComponent().Get()
	if err != nil {
		return err
	}
	defer NatsMQComponent().Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	err = encodedConn.BindSendChan(subject, sendCh)
	return err
}

// 接收主题消息，绑定管道
func (*CommonNatsChan) BindRecvChan(subject string, recvCh chan interface{}) error {
	conn, err := NatsMQComponent().Get()
	if err != nil {
		return err
	}
	defer NatsMQComponent().Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	_, err = encodedConn.BindRecvChan(subject, recvCh)
	return err
}

// 基于队列的接收操作，绑定通道。
func (*CommonNatsChan) BindRecvQueueChan(subject, queue string, recvCh chan interface{}) error {
	conn, err := NatsMQComponent().Get()
	if err != nil {
		return err
	}
	defer NatsMQComponent().Put(conn)
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	_, err = encodedConn.BindRecvQueueChan(subject, queue, recvCh)
	return err
}
