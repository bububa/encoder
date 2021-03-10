package james_stein

import (
	"github.com/bububa/encoder/category"
)

// JamesSteinRegression is a one way encoder.
// You cannot decode JamesSteinRegression values
// as some values may be encoded with the same
// numerical code.
// JamesSteinRegression is a target-based encoder.
type Regression struct {
	category.Locker
	mp map[string]float64
}

func NewRegression() *Regression {
	return &Regression{
		mp: make(map[string]float64),
	}
}

func (r *Regression) reset(size int) {
	r.mp = make(map[string]float64, size)
}

func (r *Regression) Fit(oris []string, targets []float64) error {
	l := len(oris)
	if l != len(targets) {
		return category.BoundsError
	}
	r.Lock()
	defer r.Unlock()
	r.reset(l)
	targetsSum := make(map[string]float64, l)
	targetsCount := make(map[string]float64, l)
	for idx, ori := range oris {
		targetsSum[ori] += targets[idx]
		targetsCount[ori] += 1
	}
	for k, sum := range targetsSum {
		r.mp[k] = sum / targetsCount[k]
	}
	return nil
}

func (r *Regression) Transform(ori string) (float64, error) {
	r.RLock()
	defer r.RUnlock()
	if encoded, ok := r.mp[ori]; ok {
		return encoded, nil
	}
	return 0, category.NotFoundError
}
