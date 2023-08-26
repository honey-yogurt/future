package main

import (
	"errors"
	"sync"
)

type DirectMsgHandler func() string

// MessageHandlerDistributor is a MessageHandler distributor.
type MessageHandlerDistributor struct {
	handlerMap map[string]DirectMsgHandler
	lock       sync.RWMutex
}

func newMessageHandlerDistributor() *MessageHandlerDistributor {
	return &MessageHandlerDistributor{handlerMap: make(map[string]DirectMsgHandler)}
}

func (mhd *MessageHandlerDistributor) registerHandler(
	chainId string,
	msgFlag string,
	handler DirectMsgHandler,
) error {
	mhd.lock.Lock()
	defer mhd.lock.Unlock()
	key := chainId + msgFlag
	_, exist := mhd.handlerMap[key]
	if exist {
		return errors.New("can not register handler more than once")
	}
	mhd.handlerMap[key] = handler
	return nil
}

func (mhd *MessageHandlerDistributor) cancelRegisterHandler(chainId string, msgFlag string) {
	mhd.lock.Lock()
	defer mhd.lock.Unlock()
	key := chainId + msgFlag
	delete(mhd.handlerMap, key)
}

func (mhd *MessageHandlerDistributor) handler(chainId string, msgFlag string) DirectMsgHandler {
	mhd.lock.RLock()
	defer mhd.lock.RUnlock()
	key := chainId + msgFlag
	handler, ok := mhd.handlerMap[key]
	if !ok {
		return nil
	}
	return handler
}
