package XNats

import (
	"GoWebScaffold/infras"
)

var natsMQPool *NatsPool

func SetComponent(p *NatsPool) {
	natsMQPool = p
}

func NatsMQComponent() *NatsPool {
	infras.Check(natsMQPool)
	return natsMQPool
}
