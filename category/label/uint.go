package label

import (
	"github.com/bububa/encoder/category"
)

type UintEncoder struct {
	category.Locker
	mp      map[uint]int
	inverse []uint
}

func NewUintEncoder() *UintEncoder {
	return &UintEncoder{
		mp: make(map[uint]int),
	}
}

func (e *UintEncoder) reset(size int) {
	e.mp = make(map[uint]int, size)
	e.inverse = make([]uint, size)
}

func (e *UintEncoder) encode(ori uint) int {
	if encoded, ok := e.mp[ori]; ok {
		return encoded
	}
	encoded := len(e.mp)
	e.mp[ori] = encoded
	e.inverse[encoded] = ori
	return encoded
}

func (e *UintEncoder) Decode(encoded int) (uint, error) {
	e.RLock()
	defer e.RUnlock()
	if encoded < 0 || encoded >= len(e.inverse) {
		return 0, category.BoundsError
	}
	return e.inverse[encoded], nil
}

func (e *UintEncoder) Fit(oris []uint) []int {
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

func (e *UintEncoder) Transform(oris []uint) []int {
	e.Lock()
	defer e.Unlock()
	var ret = make([]int, len(oris))
	for _, v := range oris {
		ret = append(ret, e.encode(v))
	}
	return ret
}
