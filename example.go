package customMetric

import (
	"errors"
	"time"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/customMetric".
func init() {
	modules.Register("k6/x/customMetric", new(RootModule))
}

type RootModule struct {
	customMetric *metrics.Metric
}

var _ modules.Module = &RootModule{}

func (r *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	r.customMetric = vu.InitEnv().Registry.MustNewMetric("custom_req_duration", metrics.Rate)
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

func (t *thisModule) add(x float64, tags map[string]string) error {
	state := t.vu.State()
	if state == nil {
		return errors.New("add needs to be called not in the initcontext")
	}

	timeSeries := metrics.TimeSeries{
		Metric: t.root.customMetric,
		Tags:   state.Tags.GetCurrentValues().Tags.WithTagsFromMap(state.Options.RunTags).WithTagsFromMap(tags),
		// Tags: metrics.NewRegistry().RootTagSet().WithTagsFromMap(tags),
	}

	metrics.PushIfNotDone(t.vu.Context(), t.vu.State().Samples, metrics.Sample{
		TimeSeries: timeSeries,
		Value:      x,
		Time:       time.Now(),
	})
	return nil
}
