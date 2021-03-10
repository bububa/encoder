package onehot

import (
	"github.com/bububa/encoder/category"
)

type Encoder struct {
	category.Locker
	mp      map[string]int
	inverse []string
}

func NewEncoder() *Encoder {
	return &Encoder{
		mp: make(map[string]int),
	}
}

func (e *Encoder) reset(size int) {
	e.mp = make(map[string]int, size)
	e.inverse = make([]string, size)
}

func (e *Encoder) encode(ori string) int {
	if encoded, ok := e.mp[ori]; ok {
		return encoded
	}
	encoded := len(e.mp)
	e.mp[ori] = encoded
	e.inverse[encoded] = ori
	return encoded
}

func (e *Encoder) Fit(oris []string) []int {
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

func (e *Encoder) Transform(ori string) []float32 {
	e.RLock()
	defer e.RUnlock()
	ret := make([]float32, len(e.mp))
	if idx, ok := e.mp[ori]; ok {
		ret[idx] = 1
	} else {
		for idx := range ret {
			ret[idx] = -1
		}
	}
	return ret
}

func (e *Encoder) Decode(encoded []float32) (string, error) {
	e.RLock()
	defer e.RUnlock()
	if len(encoded) > len(e.inverse) {
		return "", category.BoundsError
	}
	for idx, v := range encoded {
		if v == 1 {
			return e.inverse[idx], nil
		}
	}
	return "", category.NotFoundError
}
