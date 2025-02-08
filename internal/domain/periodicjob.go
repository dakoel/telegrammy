package domain

type PeriodicJob struct {
	Schedule string `yaml:"schedule"`
	Job      `yaml:",inline"`
}
