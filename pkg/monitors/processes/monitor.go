package processes

import (
	"github.com/signalfx/signalfx-agent/pkg/core/config"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

// Config for this monitor
type Config struct {
	config.MonitorConfig `singleInstance:"false" acceptsEndpoints:"false"`

	ProcessName string `yaml:"processName" validate:"required"`
}

func init() {
	monitors.Register(&monitorMetadata, func() interface{} { return &Monitor{} }, &Config{})
}

type Monitor struct {
}

// Configure the monitor and start collection
func (m *Monitor) Configure(conf *Config) error {
	return nil
}
