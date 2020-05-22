package natsMq

import (
	"GoWebScaffold/infras/config"
	"fmt"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

/*
NatsMq，类似于redis式的轻量级消息中间件，用于高吞吐量的应用，性能比redis高许多，但不保证可靠送达，消息发送后不管
特性：高性能（fast）、一直可用（dial tone）、极度轻量级（small footprint）、最多交付一次（fire and forget，消息发送后不管）、支持多种消息通信模型和用例场景（flexible）
应用场景：　寻址、发现、命令和控制（控制面板）、负载均衡、多路可伸缩能力、定位透明、容错等。
*/

func NatsMqPoolInit(appConf *config.AppConfig, logger *zap.Logger) *NatsPool {
	var err error
	var serverList []string
	var natsServersUrl string
	var natsMqConnPool *NatsPool
	for _, server := range appConf.MqConf.NatsMq.NatsServers {
		var natsUrl = "nats://"
		if server.Host == "" {
			natsUrl = nats.DefaultURL
		} else {
			if server.AuthSwitch {
				natsUrl += server.UserName + ":" + server.Password + "@"
			}
			natsUrl += server.Host + ":" + strconv.Itoa(server.Port)
		}
		serverList = append(serverList, natsUrl)
	}
	if len(serverList) > 1 {
		natsServersUrl = strings.Join(serverList, ",")
	} else {
		natsServersUrl = serverList[0]
	}
	fmt.Println("natsServersUrl:", natsServersUrl)

	//  nats conn 初始化连接池
	natsMqConnPool, err = NewDefaultPool(natsServersUrl, logger)
	if err != nil {
		logger.Fatal("NewDefaultPool Error:" + err.Error())
	}

	return natsMqConnPool
}
