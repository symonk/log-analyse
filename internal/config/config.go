package config

// Config encapsulates basic configuration used by
// log-analyse when it runs.  At the moment this is
// relatively basic, enhanced configurations will be
// enabled in future.
type Config struct {
	Files        []FileConfig  `yaml:"files"`
	Integrations []Integration `yaml:"integrations, omitempty"`
}

// FileConfig encapsualates the threshold for pattern
// matches before an alert or action is triggered.
type FileConfig struct {
	Glob   string `yaml:"glob"`
	Times  int    `yaml:"times"`
	Period int    `yaml:"period"`
}

// Integration is an implementation of an alerting
// mechanism
type Integration struct {
	Slack Slack `yaml:"slack, omitempty"`
	Email Email `yaml:"email, omitempty"`
}

// Slack encapsulates configurations for the slack
// notification plugin
type Slack struct {
	Webhook string `yaml:"webhook"`
}

// Email encapsulates configurations for the email
// notification plugin
type Email struct {
	To []string `yaml:"to"`
}
