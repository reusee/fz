package testkv

import "sync"

type KV struct {
	kv  sync.Map
	sem chan struct{}
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
}

func (k *KV) Get(key any) (value any) {
	k.sem <- struct{}{}
	defer func() {
		<-k.sem
	}()
	value, _ = k.kv.Load(key)
	return
}
