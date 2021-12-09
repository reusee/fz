package testkv

import (
	"sync"
	"sync/atomic"
)

type KV struct {
	kv     sync.Map
	sem    chan struct{}
	numOps int64
}

func NewKV(
	maxClients int,
) *KV {
	return &KV{
		sem: make(chan struct{}, maxClients),
	}
}

func (k *KV) Set(key any, value any) {
	k.sem <- struct{}{}
	defer func() {
		<-k.sem
	}()
	k.kv.Store(key, value)
	atomic.AddInt64(&k.numOps, 1)
}

func (k *KV) Get(key any) (value any) {
	k.sem <- struct{}{}
	defer func() {
		<-k.sem
	}()
	value, _ = k.kv.Load(key)
	atomic.AddInt64(&k.numOps, 1)
	return
}
