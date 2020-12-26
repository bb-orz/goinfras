package XRedis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

type RedisCommand struct {
	pool *redis.Pool
}

// 从Redis连接池获取一个连接
func (p *RedisCommand) GetRedisConn() redis.Conn {
	conn := p.pool.Get()
	return conn
}

// 单次执行命令的R函数,执行完命令自动关闭连接
func (p *RedisCommand) R(command string, args ...interface{}) (reply interface{}, err error) {
	conn := p.GetRedisConn()
	defer func() {
		conn.Flush()
		conn.Close()
	}()
	return conn.Do(command, args...)
}

// pipeline 串行命令，减少网络开销
// e.g.: {{"SET","name","ginger"},{"SET","key","value"},{"GET","key"}}
type CommandPipe [][]interface{}
type ReplysPipe []interface{}

func (p *RedisCommand) P(commands CommandPipe) (ReplysPipe, error) {
	conn := p.GetRedisConn()
	defer func() {
		conn.Flush()
		conn.Close()
	}()

	var err error
	var replys ReplysPipe

	for _, cp := range commands {
		if cmd, ok := cp[0].(string); ok {
			params := cp[1:]
			err = conn.Send(cmd, params...)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("commandPipe type error")
		}
	}

	err = conn.Flush()
	if err != nil {
		return nil, err
	}

	cmdCount := len(commands)
	replys = make(ReplysPipe, cmdCount)

	for i := 0; i < cmdCount; i++ {
		rs, err := conn.Receive()
		if err != nil {
			return nil, err
		}
		replys = append(replys, rs)
	}

	return replys, nil

}
