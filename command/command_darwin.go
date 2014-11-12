package command

import (
	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/metrics"
	metricsDarwin "github.com/mackerelio/mackerel-agent/metrics/darwin"
	"github.com/mackerelio/mackerel-agent/spec"
	specDarwin "github.com/mackerelio/mackerel-agent/spec/darwin"
)

func specGenerators() []spec.Generator {
	return []spec.Generator{
		&specDarwin.KernelGenerator{},
		&specDarwin.MemoryGenerator{},
		&specDarwin.CPUGenerator{},
	}
}

func interfaceGenerator() spec.Generator {
	return &specDarwin.InterfaceGenerator{}
}

func metricsGenerators(conf *config.Config) []metrics.Generator {
	generators := []metrics.Generator{
		&metricsDarwin.Loadavg5Generator{},
		&metricsDarwin.CpuusageGenerator{},
	}

	return generators
}

func pluginGenerators(conf *config.Config) []metrics.PluginGenerator {
	// TODO
	return nil
}