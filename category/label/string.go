package label

import (
	"github.com/bububa/encoder/category"
)

type StringEncoder struct {
	category.Locker
	mp      map[string]int
	inverse []string
}

func NewStringEncoder() *StringEncoder {
	return &StringEncoder{
		mp: make(map[string]int),
	}
}

func (e *StringEncoder) reset(size int) {
	e.mp = make(map[string]int, size)
	e.inverse = make([]string, size)
}

func (e *StringEncoder) encode(ori string) int {
	if encoded, ok := e.mp[ori]; ok {
		return encoded
	}
	encoded := len(e.mp)
	e.mp[ori] = encoded
	e.inverse[encoded] = ori
	return encoded
}

func (e *StringEncoder) Decode(encoded int) (string, error) {
	e.RLock()
	defer e.RUnlock()
	if encoded < 0 || encoded >= len(e.inverse) {
		return "", category.BoundsError
	}
	return e.inverse[encoded], nil
}

func (e *StringEncoder) Fit(oris []string) []int {
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

func (e *StringEncoder) Transform(oris []string) []int {
	e.Lock()
	defer e.Unlock()
	var ret = make([]int, len(oris))
	for _, v := range oris {
		ret = append(ret, e.encode(v))
	}
	return ret
}
