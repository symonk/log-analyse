package config

// Config encapsulates basic configuration used by
// log-analyse when it runs.  At the moment this is
// relatively basic, enhanced configurations will be
// enabled in future.
type Config struct {
	Files        []FileConfig
	Integrations string
}

// FileConfig encapsualates the threshold for pattern
// matches before an alert or action is triggered.
type FileConfig struct {
	glob   string `yaml:"glob"`
	times  int    `yaml:"times"`
	period int    `yaml:"period"`
}

// Integration is an implementation of an alerting
// mechanism
type Integration struct {
}
