package config

// Config encapsulates basic configuration used by
// log-analyse when it runs.  At the moment this is
// relatively basic, enhanced configurations will be
// enabled in future.
type Config struct {
	checks []Threshold
}

// Threshold encapsualates the threshold for pattern
// matches before an alert or action is triggered.
type Threshold struct {
	glob   string `yaml:"glob"`
	times  int    `yaml:"times"`
	period int    `yaml:"period"`
}
