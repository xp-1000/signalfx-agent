// +build linux

package mongodb

import (
	"context"
	"time"

	telegrafInputs "github.com/influxdata/telegraf/plugins/inputs"
	telegrafPlugin "github.com/influxdata/telegraf/plugins/inputs/mongodb"
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/core/config"
	"github.com/signalfx/signalfx-agent/pkg/monitors"

	"github.com/signalfx/signalfx-agent/pkg/monitors/telegraf/common/accumulator"
	"github.com/signalfx/signalfx-agent/pkg/monitors/telegraf/common/emitter/baseemitter"
	"github.com/signalfx/signalfx-agent/pkg/monitors/types"
	"github.com/signalfx/signalfx-agent/pkg/utils"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	monitors.Register(&monitorMetadata, func() interface{} { return &Monitor{} }, &Config{})
}

// Config for this monitor
type Config struct {
	config.MonitorConfig `yaml:",inline" singleInstance:"false" acceptsEndpoints:"true"`
	// If running as a restricted user enable this flag to prepend sudo.
	Servers             []string `yaml:"servers" validate:"required"`
	GatherClusterStatus bool     `yaml:"gatherClusterStatus"`
	GatherPerdbStats    bool     `yaml:"gatherPerdbStats"`
	GatherColStats      bool     `yaml:"gatherColStats"`
	ColStatsDbs         []string `yaml:"colStatsDbs"`
	TLSCA               string   `yaml:"tlsCa"`
	TLSCert             string   `yaml:"tlsCert"`
	TLSKey              string   `yaml:"tlsKey"`
	InsecureSkipVerify  bool     `yaml:"insecureSkipVerify"`
}

// Monitor for Utilization
type Monitor struct {
	Output types.Output
	cancel func()
	logger logrus.FieldLogger
}

type Emitter struct {
	*baseemitter.BaseEmitter
}

// Configure the monitor and kick off metric syncing
func (m *Monitor) Configure(conf *Config) (err error) {
	m.logger = logrus.WithFields(log.Fields{"monitorType": monitorType})

	plugin := telegrafInputs.Inputs["mongodb"]().(*telegrafPlugin.MongoDB)
	plugin.Servers = conf.Servers
	plugin.GatherClusterStatus = conf.GatherClusterStatus
	plugin.GatherPerdbStats = conf.GatherPerdbStats
	plugin.GatherColStats = conf.GatherColStats
	plugin.ColStatsDbs = conf.ColStatsDbs
	plugin.TLSCA = conf.TLSCA
	plugin.TLSCert = conf.TLSCert
	plugin.TLSKey = conf.TLSKey
	plugin.InsecureSkipVerify = conf.InsecureSkipVerify

	//emitter := &Emitter{baseemitter.NewEmitter(m.Output, m.logger)}
	emitter := baseemitter.NewEmitter(m.Output, m.logger)

	// don't include the telegraf_type dimension
	emitter.SetOmitOriginalMetricType(true)
	emitter.AddDatapointTransformation(func(dp *datapoint.Datapoint) error {
		m.logger.Error(dp.Metric)

		if dp.Metric == "mongodb.repl_state" {
			m.logger.Error(dp.Dimensions)
			m.logger.Error(dp.MetricType)
			m.logger.Error(dp.Value)
			dp.Value = datapoint.NewIntValue(1)
			m.logger.Error(dp.Value)
		}

		return nil
	})

	emitter.AddTag("plugin", monitorType)

	accumulator := accumulator.NewAccumulator(emitter)

	// create contexts for managing the the plugin loop
	var ctx context.Context
	ctx, m.cancel = context.WithCancel(context.Background())

	// gather metrics on the specified interval
	utils.RunOnInterval(ctx, func() {
		if err := plugin.Gather(accumulator); err != nil {
			m.logger.WithError(err).Errorf("an error occurred while gathering metrics")
		}
	}, time.Duration(conf.IntervalSeconds)*time.Second)

	return err
}

// Shutdown stops the metric sync
func (m *Monitor) Shutdown() {
	if m.cancel != nil {
		m.cancel()
	}
}
