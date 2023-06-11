// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for hostmetricsreceiver/info metrics.
type MetricsConfig struct {
	InfoNow MetricConfig `mapstructure:"info.now"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		InfoNow: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for hostmetricsreceiver/info resource attributes.
type ResourceAttributesConfig struct {
	InfoCPUNum   ResourceAttributeConfig `mapstructure:"info.cpu.num"`
	InfoHostname ResourceAttributeConfig `mapstructure:"info.hostname"`
	InfoOrg      ResourceAttributeConfig `mapstructure:"info.org"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		InfoCPUNum: ResourceAttributeConfig{
			Enabled: true,
		},
		InfoHostname: ResourceAttributeConfig{
			Enabled: true,
		},
		InfoOrg: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for hostmetricsreceiver/info metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
