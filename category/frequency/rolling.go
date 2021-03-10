package frequency

import (
	"github.com/bububa/encoder/category"
)

type RollingEncoder struct {
	category.Locker
	codes  []int
	window int
}

// NewRollingEncoder will create a codeword for every value in the list of values
// in the order of those values.
// The list of values supplied to this function should not be a unique list of categorical
// values.
// The list should contain all the individual observation values found in the dataset/sample.
func NewRollingEncoder(window int) *RollingEncoder {
	return &RollingEncoder{
		window: window,
	}
}

func (e *RollingEncoder) reset(size int) {
	e.codes = make([]int, size)
}

func (e *RollingEncoder) Fit(oris []string) {
	e.Lock()
	defer e.Unlock()
	l := len(oris)
	e.reset(l)
	mp := make(map[string]int, l)
	for idx, v := range oris {
		if idx%e.window == 0 {
			mp = make(map[string]int, l)
		}
		mp[v] += 1
		e.codes[idx] = mp[v]
	}
}

func (e *RollingEncoder) Codes() []int {
	e.RLock()
	defer e.RUnlock()
	return e.codes
}

func (e *RollingEncoder) Get(idx int) (int, error) {
	e.RLock()
	defer e.RUnlock()
	if idx < 0 || idx >= len(e.codes) {
		return 0, category.BoundsError
	}
	return e.codes[idx], nil
}
