package label

import (
	"github.com/bububa/encoder/category"
)

type Uint64Encoder struct {
	category.Locker
	mp      map[uint64]int
	inverse []uint64
}

func NewUint64Encoder() *Uint64Encoder {
	return &Uint64Encoder{
		mp: make(map[uint64]int),
	}
}

func (e *Uint64Encoder) reset(size int) {
	e.mp = make(map[uint64]int, size)
	e.inverse = make([]uint64, size)
}

func (e *Uint64Encoder) encode(ori uint64) int {
	if encoded, ok := e.mp[ori]; ok {
		return encoded
	}
	encoded := len(e.mp)
	e.mp[ori] = encoded
	e.inverse[encoded] = ori
	return encoded
}

func (e *Uint64Encoder) Decode(encoded int) (uint64, error) {
	e.RLock()
	defer e.RUnlock()
	if encoded < 0 || encoded >= len(e.inverse) {
		return 0, category.BoundsError
	}
	return e.inverse[encoded], nil
}

func (e *Uint64Encoder) Fit(oris []uint64) []int {
	e.Lock()
	defer e.Unlock()
	l := len(oris)
	e.reset(l)
	var ret = make([]int, l)
	for _, v := range oris {
		ret = append(ret, e.encode(v))
	}
	return ret
}

func (e *Uint64Encoder) Transform(oris []uint64) []int {
	e.Lock()
	defer e.Unlock()
	var ret = make([]int, len(oris))
	for _, v := range oris {
		ret = append(ret, e.encode(v))
	}
	return ret
}
