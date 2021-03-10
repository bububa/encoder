package binary

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
	l := len(e.mp)
	var ret []float32
	for ; l > 0; l /= 2 {
		ret = append(ret, 0)
	}
	if idx, ok := e.mp[ori]; ok {
		n := idx + 1
		var arr []float32
		for ; n > 0; n /= 2 {
			arr = append([]float32{float32(n % 2)}, arr...)
		}
		prefix := make([]float32, len(ret)-len(arr))
		ret = append(prefix, arr...)
	}
	return ret
}

func (e *Encoder) Decode(encoded []float64) (string, error) {
	var idx int
	for _, v := range encoded {
		n := int(v)
		if n != 1 {
			continue
		}
		idx |= 1 << n
	}
	if idx > len(e.inverse) || idx <= 0 {
		return "", category.BoundsError
	}
	return e.inverse[idx-1], nil
}
