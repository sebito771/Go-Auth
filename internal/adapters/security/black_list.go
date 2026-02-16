package security

import (
	"sync"
	"time"
)

type BlackList struct {
	store map[string]time.Time
	MU sync.RWMutex
}

func NewBlackList() *BlackList {
	bl:= &BlackList{store: make(map[string]time.Time)}
	go bl.cleanUp()
	return bl
}

func (bl *BlackList) Add(token string, ttl int64) error {
	bl.MU.Lock()
	defer bl.MU.Unlock()
	expiredAt:= time.Now().Add(time.Duration(ttl) * time.Second)
	bl.store[token] = expiredAt
	return nil
}

func (bl *BlackList) IsBlackListed(token string) (bool, error) {
	bl.MU.RLock()
	expiresAt, exist := bl.store[token]
	if !exist {
		 bl.MU.RUnlock()
		return false, nil
	 
	}
	bl.MU.RUnlock()


	 bl.MU.Lock()
	 defer bl.MU.Unlock()
	if time.Now().After(expiresAt) {
		delete(bl.store, token)
		return false, nil
	}
	return true, nil
}

func(bl *BlackList) cleanUp(){
	ticker:= time.NewTicker(time.Minute)

  for range ticker.C{
	bl.MU.Lock()
    now:= time.Now()
	for token, expiresAt:= range bl.store{
		if now.After(expiresAt){
			delete(bl.store,token)
			
		}	
	}
	bl.MU.Unlock()
  }

}