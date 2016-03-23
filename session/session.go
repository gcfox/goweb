package session

import (
	"sync"
	"fmt"
)

var provides = make(map[string]Provider)

type Manage struct {
    cookieName string
    lock sync.Mutex
    provider Provider
    maxlifetime int64
}

func NewManage(provideName, cookieName string, maxlifetime int64) (*Manage, error) {
    provider, ok := provides[provideName]
    
    if !ok {
        return nil, fmt.Errorf("session: unknow provider %q (forgoten import?)", provideName)
    }
    
    return &Manage{provider:provider, cookieName:cookieName, maxlifetime: maxlifetime}, nil
    
}

type Provider interface {
    SessionInint(sid string) (Session, error)
    SessionRead(sid string) (Session, error)
    SessionDestroy(sid string) error
    SessionGC(maxLifeTime int64)
}

type Session interface {
    Set(key, value interface{}) error
    Get(key interface{}) interface{}
    Delete(key interface{}) error
    SessionId() string
}