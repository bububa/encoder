package category

import (
	"sync"
)

type Locker struct {
	lock *sync.RWMutex
}

func (l *Locker) InitLocker() {
	if l.lock != nil {
		return
	}
	l.lock = new(sync.RWMutex)
}

func (l *Locker) Lock() {
	if l.lock == nil {
		return
	}
	l.lock.Lock()
}

func (l *Locker) Unlock() {
	if l.lock == nil {
		return
	}
	l.lock.Unlock()
}

func (l *Locker) RLock() {
	if l.lock == nil {
		return
	}
	l.lock.RLock()
}

func (l *Locker) RUnlock() {
	if l.lock == nil {
		return
	}
	l.lock.RUnlock()
}
