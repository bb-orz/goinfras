package XNats

import (
	"errors"
	"github.com/nats-io/nats.go"
	"reflect"
)

type commonNatsPubSub struct {
	pool *NatsPool
}

/*
发布消息到一个主题
@param subject string  发布主题
@param msg interface{} 发布的消息
*/
func (c *commonNatsPubSub) Publish(subject string, msg interface{}) error {
	conn, err := c.pool.Get()
	if err != nil {
		return err
	}
	defer c.pool.Put(conn)

	switch reflect.TypeOf(msg).Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Ptr:
		return c.publishEncodedJson(conn, subject, msg)
	case reflect.String:
		return c.publishString(conn, subject, msg.(string))
	default:
		return errors.New("Message Type Illegal")
	}
}

// 发送字符串消息类型，自动转[]byte
func (*commonNatsPubSub) publishString(conn *nats.Conn, subject string, msg string) error {
	return conn.Publish(subject, []byte(msg))
}

// 发送需要编码的go type消息类型
func (*commonNatsPubSub) publishEncodedJson(conn *nats.Conn, subject string, msg interface{}) error {
	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	return encodedConn.Publish(subject, msg)
}

/*
订阅并异步接收主题数据
@param subject string  订阅主题
@param cb nats.Handler 订阅消息处理函数
For example：

一、关于订阅主题：
//1.主题全称
subject1 := "testSubject"

//2."*"通配符匹配相应位置的字符主题
wildcardSubject1 := "foo.*.baz"

//3.">"通配符匹配任何尾部任意长度的字符主题
// E.g. 'foo.>' will match 'foo.bar', 'foo.bar.baz', 'foo.foo.bar.bax.22'
wildcardSubject1 := "foo.>"

二、关于消息处理函数：
handler := func(m *Msg)
handler := func(p *person)
handler := func(subject string, o *obj)
handler := func(subject, reply string, o *obj)   for Request() reply
*/

func (c *commonNatsPubSub) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	conn, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	defer c.pool.Put(conn)

	s, err := conn.Subscribe(subject, handler)
	if err != nil {
		return nil, err
	}

	return s, nil
}

/*
接收已编码消息的订阅，用于订阅发布go类型数据消息的主题
除接收处理函数不同，其他都一样，请自定义接收消息的数据类型，消息用json编码发送
*/
type EncodedMsgHandler func(subject string, goDataMsg interface{})

func (c *commonNatsPubSub) SubscribeForEncodedMsg(subject string, handler EncodedMsgHandler) (*nats.Subscription, error) {
	conn, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	defer c.pool.Put(conn)

	encodedConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	s, err := encodedConn.Subscribe(subject, handler)
	if err != nil {
		return nil, err
	}

	return s, nil
}

/*
取消订阅一个或多个主题
param subject/subjects string 已订阅的主题
*/
func (c *commonNatsPubSub) Unsubscribe(subscribers ...*nats.Subscription) error {
	for _, subscriber := range subscribers {
		var err error
		err = subscriber.Unsubscribe()
		if err != nil {
			return err
		}
	}
	return nil
}
