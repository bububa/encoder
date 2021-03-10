package frequency

import (
	"github.com/bububa/encoder/category"
)

type Encoder struct {
	category.Locker
	mp      map[string]int
	inverse map[int]string
}

// NewEncoder will return a frequency encoder
// with the given values encoded.
func NewEncoder() *Encoder {
	return &Encoder{
		mp:      make(map[string]int),
		inverse: make(map[int]string),
	}
}

func (e *Encoder) reset(size int) {
	e.mp = make(map[string]int, size)
	e.inverse = make(map[int]string, size)
}

func (e *Encoder) Fit(oris []string) []int {
	e.Lock()
	defer e.Unlock()
	l := len(oris)
	e.reset(l)
	var ret = make([]int, l)
	for _, v := range oris {
		e.mp[v] += 1
	}
	for ori, n := range e.mp {
		e.inverse[n] = ori
		ret = append(ret, n)
	}
	return ret
}

func (e *Encoder) Transform(ori string) float32 {
	e.RLock()
	defer e.RUnlock()
	if v, ok := e.mp[ori]; ok {
		return float32(v)
	}
	return -1
}

func (e *Encoder) Decode(encoded float32) (string, error) {
	e.RLock()
	defer e.RUnlock()
	if ori, ok := e.inverse[int(encoded)]; ok {
		return ori, nil
	}
	return "", category.NotFoundError
}
