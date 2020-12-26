# XRedis Starter

> 基于 github.com/gomodule/redigo 包

### Redis Documentation

> Documentation  请参阅文档：http://doc.redisfans.com

下面是一些常见的Redis 键值部分的命令的用法和释义

--------------------------------------------------------------------
##### [SET Command]
> SET key value [EX seconds] [PX milliseconds] [NX|XX]

将字符串值 value 关联到 key 。
如果 key 已经持有其他值， SET 就覆写旧值，无视类型。
对于某个原本带有生存时间（TTL）的键来说， 当 SET 命令成功在这个键上执行时， 这个键原有的 TTL 将被清除。

- 可选参数：
从 Redis 2.6.12 版本开始， SET 命令的行为可以通过一系列参数来修改：
    - EX second ：设置键的过期时间为 second 秒。 SET key value EX second 效果等同于 SETEX key second value 。
    - PX millisecond ：设置键的过期时间为 millisecond 毫秒。 SET key value PX millisecond 效果等同于 PSETEX key millisecond value 。
    - NX ：只在键不存在时，才对键进行设置操作。 SET key value NX 效果等同于 SETNX key value 。
    - XX ：只在键已经存在时，才对键进行设置操作。

- 返回值：
    - 在 Redis 2.6.12 版本以前， SET 命令总是返回 OK 。
    - 从 Redis 2.6.12 版本开始， SET 在设置操作成功完成时，才返回 OK 。

如果设置了 NX 或者 XX ，但因为条件没达到而造成设置操作未执行，那么命令返回空批量回复（NULL Bulk Reply）。


--------------------------------------------------------------------
##### [EXPIRE Command]
> EXPIRE key seconds

为给定 key 设置生存时间，当 key 过期时(生存时间为 0 )，它会被自动删除。
在 Redis 中，带有生存时间的 key 被称为『易失的』(volatile)。
生存时间可以通过使用 DEL 命令来删除整个 key 来移除，或者被 SET 和 GETSET 命令覆写(overwrite)，这意味着，如果一个命令只是修改(alter)一个带生存时间的 key 的值而不是用一个新的 key 值来代替(replace)它的话，那么生存时间不会被改变。
比如说，对一个 key 执行 INCR 命令，对一个列表进行 LPUSH 命令，或者对一个哈希表执行 HSET 命令，这类操作都不会修改 key 本身的生存时间。
另一方面，如果使用 RENAME 对一个 key 进行改名，那么改名后的 key 的生存时间和改名前一样。

RENAME 命令的另一种可能是，尝试将一个带生存时间的 key 改名成另一个带生存时间的 another_key ，这时旧的 another_key (以及它的生存时间)会被删除，然后旧的 key 会改名为 another_key ，因此，新的 another_key 的生存时间也和原本的 key 一样。
使用 PERSIST 命令可以在不删除 key 的情况下，移除 key 的生存时间，让 key 重新成为一个『持久的』(persistent) key 。

- 更新生存时间
可以对一个已经带有生存时间的 key 执行 EXPIRE 命令，新指定的生存时间会取代旧的生存时间。

- 过期时间的精确度
    - 在 Redis 2.4 版本中，过期时间的延迟在 1 秒钟之内 —— 也即是，就算 key 已经过期，但它还是可能在过期之后一秒钟之内被访问到，而在新的 Redis 2.6 版本中，延迟被降低到 1 毫秒之内。
    - Redis 2.1.3 之前的不同之处在 Redis 2.1.3 之前的版本中，修改一个带有生存时间的 key 会导致整个 key 被删除，这一行为是受当时复制(replication)层的限制而作出的，现在这一限制已经被修复。

- 返回值：
设置成功返回 1 。
当 key 不存在或者不能为 key 设置生存时间时(比如在低于 2.1.3 版本的 Redis 中你尝试更新 key 的生存时间)，返回 0 。

--------------------------------------------------------------------

##### [EXISTS Command]
> EXISTS key

检查给定 key 是否存在。

- 返回值：
若 key 存在，返回 1 ，否则返回 0 。

--------------------------------------------------------------------

##### [TTL Command]
> TTL key

以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)。

- 返回值：
    - 当 key 不存在时，返回 -2 。
    - 当 key 存在但没有设置剩余生存时间时，返回 -1 。
    - 否则，以秒为单位，返回 key 的剩余生存时间。

--------------------------------------------------------------------

##### [GET Command]
> GET key

返回 key 所关联的字符串值。如果 key 不存在那么返回特殊值 nil 。假如 key 储存的值不是字符串类型，返回一个错误，因为 GET 只能用于处理字符串值。

- 返回值：
    - 当 key 不存在时，返回 nil ，否则，返回 key 的值。
    - 如果 key 不是字符串类型，那么返回一个错误。


--------------------------------------------------------------------

##### [INCRBY Command]
> INCRBY key increment

将 key 所储存的值加上增量 increment 。
如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCRBY 命令。
如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
本操作的值限制在 64 位(bit)有符号数字表示之内。
关于递增(increment) / 递减(decrement)操作的更多信息，参见 INCR 命令。

- 返回值：
    - 加上 increment 之后， key 的值。

--------------------------------------------------------------------

##### [DECRBY Command]
> DECRBY key decrement

将 key 所储存的值减去减量 decrement 。
如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECRBY 操作。
如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
本操作的值限制在 64 位(bit)有符号数字表示之内。
关于更多递增(increment) / 递减(decrement)操作的更多信息，请参见 INCR 命令。

- 返回值：
    - 减去 decrement 之后， key 的值。

--------------------------------------------------------------------

##### [DEL Command]

DEL key [key ...]

删除给定的一个或多个 key 。
不存在的 key 会被忽略。

返回值：
被删除 key 的数量。

--------------------------------------------------------------------

### XRedis Starter Usage
```
goinfras.RegisterStarter(XRedis.NewStarter())

```

### XRedis Config Setting

```
DbHost      string // 主机地址
DbPort      int    // 主机端口
DbAuth      bool   // 是否开启鉴权
DbPasswd    string // 鉴权密码
MaxActive   int64  // 最大活动链接数。0为无限
MaxIdle     int64  // 最大闲置链接数，0为无限
IdleTimeout int64  // 闲置链接超时时间
```

### X  Usage

```

```