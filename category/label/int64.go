package label

import (
	"github.com/bububa/encoder/category"
)

type Int64Encoder struct {
	category.Locker
	mp      map[int64]int
	inverse []int64
}

func NewInt64Encoder() *Int64Encoder {
	return &Int64Encoder{
		mp: make(map[int64]int),
	}
}

func (e *Int64Encoder) reset(size int) {
	e.mp = make(map[int64]int, size)
	e.inverse = make([]int64, size)
}

func (e *Int64Encoder) encode(ori int64) int {
	if encoded, ok := e.mp[ori]; ok {
		return encoded
	}
	encoded := len(e.mp)
	e.mp[ori] = encoded
	e.inverse[encoded] = ori
	return encoded
}

func (e *Int64Encoder) Decode(encoded int) (int64, error) {
	e.RLock()
	defer e.RUnlock()
	if encoded < 0 || encoded >= len(e.inverse) {
		return 0, category.BoundsError
	}
	return e.inverse[encoded], nil
}

func (e *Int64Encoder) Fit(oris []int64) []int {
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

func (e *Int64Encoder) Transform(oris []int64) []int {
	e.Lock()
	defer e.Unlock()
	var ret = make([]int, len(oris))
	for _, v := range oris {
		ret = append(ret, e.encode(v))
	}
	return ret
}
