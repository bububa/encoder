package james_stein

import (
	"github.com/bububa/encoder/category"
)

// JamesSteinClassification is a one way encoder.
// You cannot decode JamesSteinClassification values
// as some values may be encoded with the same
// numerical code.
// JamesSteinClassification is a target-based encoder.
type Classification struct {
	category.Locker
	codes []float64
}

func NewClassification() *Classification {
	return &Classification{}
}

func (c *Classification) reset(size int) {
	c.codes = make([]float64, size)
}

func (c *Classification) Fit(oris []string, targets []string) error {
	l := len(oris)
	if l != len(targets) {
		return category.BoundsError
	}
	c.Lock()
	defer c.Unlock()
	c.reset(l)
	groupCounts := make(map[string]float64, l)
	classCounts := make(map[string]float64, l)
	groupClassCounts := make(map[string]map[string]float64, l)
	for idx, ori := range oris {
		class := targets[idx]
		groupCounts[ori] += 1
		classCounts[class] += 1
		if _, ok := groupClassCounts[ori]; !ok {
			groupClassCounts[ori] = make(map[string]float64)
		}
		groupClassCounts[ori][class] += 1
	}
	lf64 := float64(l)
	groupClassBValues := make(map[string]map[string]float64, l)
	for group, classCounts := range groupClassCounts {
		groupCount := groupCounts[group]
		for class, count := range classCounts {
			classCount := classCounts[class]
			groupClassPercentage := count / classCount
			classPercentage := classCount / lf64
			groupClassValue := (groupClassPercentage * (1 - groupClassPercentage)) / groupCount
			classValue := (classPercentage * (1 - classPercentage)) / lf64
			groupClassBValues[group][class] = groupClassValue / (groupClassValue + classValue)
		}
	}
	for idx, ori := range oris {
		class := targets[idx]
		c.codes[idx] = groupClassBValues[ori][class]
	}
	return nil
}

func (c *Classification) Codes() []float64 {
	c.RLock()
	defer c.RUnlock()
	return c.codes
}

func (c *Classification) Get(idx int) (float64, error) {
	c.RLock()
	defer c.RUnlock()
	if idx < 0 || idx >= len(c.codes) {
		return 0, category.BoundsError
	}
	return c.codes[idx], nil
}
