package testAsyncMap

import (
	"errors"
	"sync"
)

type CasheMap interface {
	Read(string) (string, error)
	Write(string, string) error
	Delete(string) error
}

type asyncMap struct {
	mux      *sync.RWMutex
	asyncMap map[string]string
}

func NewAsyncMap() CasheMap {
	return &asyncMap{
		&sync.RWMutex{},
		map[string]string{},
	}
}

func (a *asyncMap) Read(key string) (string, error) {
	a.mux.RLock()
	res, ok := a.asyncMap[key]
	if ok {
		a.mux.RUnlock()
		return res, nil
	}
	a.mux.RUnlock()
	return "", errors.New("data not found")
}

func (a *asyncMap) Write(key, data string) error {
	a.mux.Lock()
	a.asyncMap[key] = data
	a.mux.Unlock()
	return nil
}

func (a *asyncMap) Delete(key string) error {
	a.mux.Lock()
	delete(a.asyncMap, key)
	a.mux.Unlock()
	return nil
}
