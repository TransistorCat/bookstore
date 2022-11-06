package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	provides    = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store:Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("store:Register called twice for provide" + name)
	}
	provides[name] = p
} //添加（注册）

func New(provideName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := provides[provideName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store:unknown provider %s", provideName)
	}
	return p, nil
} //新建
