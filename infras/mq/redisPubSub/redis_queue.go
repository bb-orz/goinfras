package redisPubSub

import redigo "github.com/garyburd/redigo/redis"

var RedisQueue *RedisList

// 实现GingerQueue接口的entity,使用redis list数据结构
type RedisList struct {
	conn redigo.Conn
}

// 获取一个redis连接的 list队列
func GetRedisList(conn redigo.Conn) *RedisList {
	list := new(RedisList)
	list.conn = conn
	return list
}

// 加入队列元素
func (q *RedisList) Push(keyName string, msg interface{}) (int, error) {
	size, err := redigo.Int(q.conn.Do("LPUSH", keyName, msg))
	return size, err
}

// 单次弹出队列元素
func (q *RedisList) Pop(keyName string) (interface{}, error) {
	element, err := redigo.String(q.conn.Do("RPOP", keyName))
	return element, err
}

func (q *RedisList) Size(keyName string) (int, error) {
	size, err := redigo.Int(q.conn.Do("LLEN", keyName))
	return size, err
}

// 循环阻塞，即时弹出队列元素
// func (q *RedisList) Bpop(keyName string,reply chan interface{}) {
//
// 	loop:
// 		for {
// 			element, err := redigo.String(q.conn.Do("BRPOP", keyName))
// 			if err != nil {
// 				break loop
// 			}
//
// 			reply <- element
// 		}
//
// }
