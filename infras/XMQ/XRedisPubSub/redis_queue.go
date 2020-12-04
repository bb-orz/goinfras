package XRedisPubSub

import redigo "github.com/garyburd/redigo/redis"

// 使用redis list数据结构
type redisList struct {
	conn redigo.Conn
}

// 加入队列元素
func (q *redisList) Push(keyName string, msg interface{}) (int, error) {
	size, err := redigo.Int(q.conn.Do("LPUSH", keyName, msg))
	return size, err
}

// 单次弹出队列元素
func (q *redisList) Pop(keyName string) (interface{}, error) {
	element, err := redigo.String(q.conn.Do("RPOP", keyName))
	return element, err
}

func (q *redisList) Size(keyName string) (int, error) {
	size, err := redigo.Int(q.conn.Do("LLEN", keyName))
	return size, err
}

func (q *redisList) Close() error {
	return q.conn.Close()
}

// 循环阻塞，即时弹出队列元素
// func (q *redisList) Bpop(keyName string,reply chan interface{}) {
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
