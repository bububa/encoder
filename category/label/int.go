package label

import (
	"github.com/bububa/encoder/category"
)

type IntEncoder struct {
	category.Locker
	mp      map[int]int
	inverse []int
}

func NewIntEncoder() *IntEncoder {
	return &IntEncoder{
		mp: make(map[int]int),
	}
}
func (e *IntEncoder) reset(size int) {
	e.mp = make(map[int]int, size)
	e.inverse = make([]int, size)
}

func (e *IntEncoder) encode(ori int) int {
	if encoded, ok := e.mp[ori]; ok {
		return encoded
	}
	encoded := len(e.mp)
	e.mp[ori] = encoded
	e.inverse[encoded] = ori
	return encoded
}

func (e *IntEncoder) Decode(encoded int) (int, error) {
	e.RLock()
	defer e.RUnlock()
	if encoded < 0 || encoded >= len(e.inverse) {
		return 0, category.BoundsError
	}
	return e.inverse[encoded], nil
}

func (e *IntEncoder) Fit(oris []int) []int {
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

func (e *IntEncoder) Transform(oris []int) []int {
	e.Lock()
	defer e.Unlock()
	var ret = make([]int, len(oris))
	for _, v := range oris {
		ret = append(ret, e.encode(v))
	}
	return ret
}
