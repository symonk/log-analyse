package config

// Config encapsulates basic configuration used by
// log-analyse when it runs.  At the moment this is
// relatively basic, enhanced configurations will be
// enabled in future.
type Config struct {
	Files        []FileConfig  `yaml:"files"`
	Integrations []Integration `yaml:"integrations, omitempty"`
}

// Globs returns the configured glob patterns defined in
// the configuration file.
func (c Config) Globs() []string {
	globs := make([]string, 0, len(c.Files))
	for _, file := range c.Files {
		globs = append(globs, file.Glob)
	}
	return globs
}

// FileConfig encapsualates the threshold for pattern
// matches before an alert or action is triggered.
type FileConfig struct {
	Glob        string      `yaml:"glob" validate:"required"`
	Threshold   Threshold   `yaml:"Threshold" validate:"required"`
	Integration Integration `yaml:"Integration, omitempty"`
}

// Threshold encapsulates the configuration for each defined
// glob pattern in the config
type Threshold struct {
	Hits     int      `yaml:"hits"`
	Period   string   `yaml:"period"`
	Patterns []string `yaml:"patterns"`
	Mode     string   `yaml:"mode"`
	Notify   string   `yaml:"notify, omitempty"`
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
