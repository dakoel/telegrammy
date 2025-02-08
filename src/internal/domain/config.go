package domain

type Config struct {
	ShellPath    string        `yaml:"shellPath,omitempty"`
	ResponseJobs []ResponseJob `yaml:"responseJobs"`
	PeriodicJobs []PeriodicJob `yaml:"periodicJobs"`
	PollInterval int           `yaml:"pollInterval,omitempty"`
}
