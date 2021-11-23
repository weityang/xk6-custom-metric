package customMetric

import (
	"errors"
	"time"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/stats"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/customMetric".
func init() {
	modules.Register("k6/x/customMetric", new(RootModule))
}

type RootModule struct {
	customMetric *stats.Metric
}

var _ modules.Module = &RootModule{}

func (r *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	r.customMetric = vu.InitEnv().Registry.MustNewMetric("coolname", stats.Trend, stats.Time)
	return &thisModule{
		vu:   vu,
		root: r,
	}
}

// thisModule is just an example on how to work with custom metrics
type thisModule struct {
	vu   modules.VU
	root *RootModule
}

func (t *thisModule) Exports() modules.Exports {
	return modules.Exports{
		Default: map[string]interface{}{
			"add": t.add,
		},
	}
}

func (t *thisModule) add(x float64) error {
	if t.vu.State() == nil {
		return errors.New("add needs to be called not in the initcontext")
	}
	stats.PushIfNotDone(t.vu.Context(), t.vu.State().Samples, stats.Sample{
		Metric: t.root.customMetric,
		Value:  x,
		Time:   time.Now(),
		Tags:   stats.NewSampleTags(t.vu.State().CloneTags()),
	})
	return nil
}
